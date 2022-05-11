package cuetojson

import (
	"cuelang.org/go/cue"
	"fmt"
)

func PrintAsJSON(value *cue.Value) {
	fmt.Printf("JSON Format:\n%v\n", value)
}

func PrintAsCUE(value *cue.Value) {
	fmt.Printf("CUE Format:\n%#v\n", value)
}
