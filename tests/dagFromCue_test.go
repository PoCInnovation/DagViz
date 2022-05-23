package tests

import (
	"cuelang.org/go/cue/cuecontext"
	"dagviz/cuetojson"
	"dagviz/dag"
	"testing"
)

func TestDagFromCue(t *testing.T) {
	args := []string{"dagviz", "samplesTests/test.cue"}
	context := cuecontext.New()

	programs, err := cuetojson.LoadFile(context, args[1])

	if err != nil {
		return
	}

	root := dag.CreateDag("dag")
	cuetojson.AppendValuesToDag(root, programs)

	if root.Members[0].Value != "hello" {
		t.Errorf("Expected hello, got %s", root.Members[0].Value)
	}

	if root.Members[0].Links[0].Value != "foo = bar" {
		t.Errorf("Expected foo, got %s", root.Members[0].Links[0].Value)
	}

	if root.Members[0].Links[1].Value != "fee" {
		t.Errorf("Expected fee, got %s", root.Members[0].Links[0].Value)
	}

	if root.Members[0].Links[1].Links[0].Value != "salut" {
		t.Errorf("Expected salut, got %s", root.Members[0].Links[0].Value)
	}

	if root.Members[0].Links[1].Links[0].Links[0].Value != "bonsoir = au revoir" {
		t.Errorf("Expected \"bonsoir = au revoir\", got %s", root.Members[0].Links[0].Value)
	}

	if root.Members[1].Value != "info = bonsoir" {
		t.Errorf("Expected \"info = bonsoir\", got %s", root.Members[1].Value)
	}
}
