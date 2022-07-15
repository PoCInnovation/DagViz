package main

import (
	"cuelang.org/go/cue/cuecontext"
	cueToDag2 "dagviz/go/cueToDag"
	"flag"
	"fmt"
	"os"
)

var (
	jsonOutput bool
	CueDir     string
)

func init() {
	flag.BoolVar(&jsonOutput, "j", false, "Prints the DAG in JSON format")
	flag.StringVar(&CueDir, "c", "", "The directory where the CUE files are located")
	flag.Parse()
}

func main() {
	context := cuecontext.New()

	if len(CueDir) <= 0 {
		fmt.Println("Error: A CUE file is required")
		os.Exit(1)
	}

	programs, err := cueToDag2.LoadFile(context, CueDir, nil)

	if err != nil {
		return
	}

	infos := cueToDag2.ExtractInfos(programs)
	root := cueToDag2.CreateCueDag("dag")

	cueToDag2.LinkDefinitions(infos, &root)

	if jsonOutput == true {
		cueToDag2.PrintJson(root.Members[0].Links)
	} else {
		root.PrintDag("-->")
	}
}
