package cueToDag

import (
	"cuelang.org/go/cue/cuecontext"
	"testing"
)

func TestPrintJson(t *testing.T) {
	CueDir := "../tests/package_test"
	context := cuecontext.New()
	programs, err := LoadFile(context, CueDir, nil)

	if err != nil {
		return
	}

	infos := ExtractInfos(programs)
	root := CreateCueDag("dag")
	LinkDefinitions(infos, &root)
	root.PrintDag("-->")
}
