package main

import (
	"cuelang.org/go/cue/cuecontext"
	"dagviz/cueToDag"
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

	programs, err := cueToDag.LoadFile(context, args[1], nil)

	if err != nil {
		return
	}

	infos := cueToDag.ExtractInfos(programs)
	root := cueToDag.CreateCueDag("dag")

	cueToDag.LinkDefinitions(infos, &root)
	root.PrintDag(" --> ")
}
