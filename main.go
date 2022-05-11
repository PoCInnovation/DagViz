package main

import (
	"cuelang.org/go/cue/cuecontext"
	"cuetojson/cuetojson"
	"fmt"
	"os"
)

func main() {
	args := os.Args
	context := cuecontext.New()

	if len(args) != 2 {
		fmt.Println("Error: A CUE file is required")
		os.Exit(1)
	}

	value, err := cuetojson.LoadFile(context, args[1])

	if err != nil {
		return
	}

	cuetojson.PrintAsJSON(value)
	cuetojson.PrintAsCUE(value)
}
