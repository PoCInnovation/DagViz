package main

import (
	"cuelang.org/go/cue/cuecontext"
	"dagviz/dagviz"
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

	programs, err := dagviz.LoadFile(context, args[1], nil)

	if err != nil {
		return
	}

	infos := dagviz.ExtractInfos(programs)
	root := dagviz.CreateCueDag("dag")

	dagviz.LinkDefinitions(infos, &root)
	root.PrintDag(" --> ")
}
