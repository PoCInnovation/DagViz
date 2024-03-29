package main

import (
	"flag"
	"fmt"
	"os"

	"cuelang.org/go/cue/cuecontext"
	"github.com/PoCInnovation/DagViz/cueToDag"
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
	err := run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	context := cuecontext.New()

	if len(CueDir) <= 0 {
		return fmt.Errorf("a CUE file is required")
	}

	programs, err := cueToDag.LoadFile(context, CueDir, nil)
	if err != nil {
		return fmt.Errorf("load file error: %w", err)
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

	return nil
}
