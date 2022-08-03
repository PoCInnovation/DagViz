package cueToDag

import (
	"strings"
	"testing"

	"cuelang.org/go/cue/cuecontext"
	"github.com/stretchr/testify/assert"
)

func TestCreateDag(t *testing.T) {
	cueRoot := CreateCueDag("Test Dag")

	assert.Equal(t, cueRoot.Root.Name, "Test Dag", "Root name are different")
	assert.Equal(t, len(cueRoot.Root.Members), 0, "Root members are different")
}

func TestBuildError(t *testing.T) {
	context := cuecontext.New()

	_, err := LoadFile(context, "../samples/build_error", nil)

	assert.Equal(t, err != nil, true, "No error thrown")
}

func TestNoDefinitions(t *testing.T) {
	context := cuecontext.New()

	programs, err := LoadFile(context, "../samples/no_def", nil)
	if err != nil {
		t.Fail()
	}

	infos := ExtractInfos(programs)
	root := CreateCueDag("dag")

	LinkDefinitions(infos, &root)
	assert.Equal(t, len(root.Members), 1, "Root members are different")
	assert.Equal(t, len(root.Members[0].Links), 0, "Links are different")
}

func TestDefinitions(t *testing.T) {
	context := cuecontext.New()

	programs, err := LoadFile(context, "../samples/def", nil)
	if err != nil {
		t.Fail()
	}

	infos := ExtractInfos(programs)
	root := CreateCueDag("dag")

	LinkDefinitions(infos, &root)
	assert.Equal(t, len(root.Members), 1, "Root members are different")
	assert.Equal(t, len(root.Members[0].Links), 1, "Links are different")

	link := root.Members[0].Links[0]
	definition := NodeDefinition{
		name: "#Info",
		file: "/cue.mod/pkg/test.fr/info/info.cue",
		def:  "#Info: {\n   name: \"info\"\n   description: \"program information\"\n   version: 2\n}",
	}
	assert.Equal(t, link.Value, definition, "Definition is different")
}

func TestCueInfos(t *testing.T) {
	context := cuecontext.New()

	programs, err := LoadFile(context, "../samples/infos", nil)
	if err != nil {
		t.Fail()
	}

	infos := ExtractInfos(programs)
	assert.Equal(t, len(infos), 1, "Infos length is different")
	assert.Equal(t, len(infos[0].Tags), 0, "Tags length is different")
	assert.Equal(t, strings.Contains(infos[0].Root, "samples/infos"), true, "Root does not include path")
	assert.Equal(t, infos[0].Module, "", "Module is different")
	assert.Equal(t, infos[0].Package, "testing", "Package is different")
	assert.Equal(t, len(infos[0].Files), 2, "Files length is different")
	assert.Equal(t, len(infos[0].IgnoredFiles), 0, "There should be no ignored files")
	assert.Equal(t, len(infos[0].InvalidFiles), 0, "There should be no invalid files")
	assert.Equal(t, len(infos[0].OrphanedFiles), 0, "There should be no orphaned files")
}
