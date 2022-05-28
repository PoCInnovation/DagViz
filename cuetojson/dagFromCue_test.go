package cuetojson

import (
	"cuelang.org/go/cue/cuecontext"
	"dagviz/dag"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDagFromCue(t *testing.T) {
	args := []string{"dagviz", "../samples/test.cue"}
	context := cuecontext.New()

	programs, err := LoadFile(context, args[1])

	if err != nil {
		return
	}

	root := dag.CreateDag("dag")
	AppendValuesToDag(root, programs)

	assert.Equalf(t, root.Members[0].Value == "hello", true, "Expected hello, got %s", root.Members[0].Value)
	assert.Equalf(t, root.Members[0].Links[0].Value == "foo = bar", true, "Expected foo, got %s", root.Members[0].Links[0].Value)
	assert.Equalf(t, root.Members[0].Links[1].Value == "fee", true, "Expected fee, got %s", root.Members[0].Links[0].Value)
	assert.Equalf(t, root.Members[0].Links[1].Links[0].Value == "salut", true, "Expected salut, got %s", root.Members[0].Links[0].Value)
	assert.Equalf(t, root.Members[0].Links[1].Links[0].Links[0].Value == "bonsoir = au revoir", true, "Expected \"bonsoir = au revoir\", got %s", root.Members[0].Links[0].Value)
	assert.Equalf(t, root.Members[1].Value == "info = bonsoir", true, "Expected \"info = bonsoir\", got %s", root.Members[1].Value)
}
