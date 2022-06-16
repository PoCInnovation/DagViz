package main

import (
	"cuelang.org/go/cue/cuecontext"
	"dagviz/cuetojson"
	"dagviz/dag"
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

	programs, err := cuetojson.LoadFile(context, args[1], nil)

	if err != nil {
		return
	}

	root := dag.CreateDag("dag")
	cuetojson.AppendValuesToDag(root, programs)

	infos := cuetojson.ExtractInfos(programs)
	cuetojson.PrintAsJSON(infos)
}
