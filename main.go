package main

import (
	"cuelang.org/go/cue/cuecontext"
	"dagviz/cuetojson"
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

	infos := cuetojson.ExtractInfos(programs)
	root := cuetojson.CreateCueDag("dag")

	cuetojson.LinkDefinitions(infos, &root)
	root.PrintCueDag(" --> ")
}
