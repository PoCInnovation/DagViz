package main

import (
	"cuelang.org/go/cue/cuecontext"
	"dagviz/cueToDag"
	"flag"
	"fmt"
	"os"
)

var (
	jsonOutput bool
	CueDir     string
)

func init() {
	flag.BoolVar(&jsonOutput, "json", false, "Prints the DAG in JSON format")
	flag.StringVar(&CueDir, "c", "", "The directory where the CUE files are located")
	flag.Parse()
}

func main() {
	context := cuecontext.New()

	if len(CueDir) <= 0 {
		fmt.Println("Error: A CUE file is required")
		os.Exit(1)
	}

	programs, err := cueToDag.LoadFile(context, CueDir, nil)

	if err != nil {
		return
	}

	infos := cueToDag.ExtractInfos(programs)
	root := cueToDag.CreateCueDag("dag")

	cueToDag.LinkDefinitions(infos, &root)

	if jsonOutput {
		fmt.Printf("{\"file\": \"%s\",\"dag\":", CueDir)
		cueToDag.PrintJson(root.Members[0].Links)
		fmt.Printf("}")
	} else {
		root.PrintDag("-->")
	}
}
