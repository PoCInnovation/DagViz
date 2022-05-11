package main

import (
	"cuelang.org/go/cue/cuecontext"
	"cuetojson/cuetojson"
	"fmt"
	"strings"
	"testing"
)

func TestSimpleCue(t *testing.T) {
	file := "samples/simple.cue"
	context := cuecontext.New()
	result, err := cuetojson.LoadFile(context, file)

	lenExpected := 1
	valueExpected := "{\n        hello: \"world\"\n}"

	if err != nil {
		t.Errorf("An error occured: %s", err)
	} else if len(result) != lenExpected {
		t.Errorf("Number of CUE instances is different (%d) from expected (%d)", len(result), lenExpected)
	}

	value := fmt.Sprintf("%v", result[0].Values)

	if strings.Compare(value, valueExpected) == 0 {
		t.Errorf("CUE values are different (%s) from expected (%s)", value, valueExpected)
	}
}
