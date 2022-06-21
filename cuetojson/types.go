package cuetojson

import (
	"cuelang.org/go/cue"
	"cuelang.org/go/cue/build"
	"dagviz/dag"
	"fmt"
)

type CueProgram struct {
	Instance *build.Instance
	Values   *cue.Value
}

type CueInfos struct {
	Tags          []string
	Root          string
	Module        string
	Package       string
	Dependencies  []string
	Directory     string
	Files         []string
	BuildFiles    []string
	InvalidFiles  []string
	IgnoredFiles  []string
	OrphanedFiles []string
	Imports       []CueInfos
}

type Definition struct {
	file string
	def  string
}

type NodeDefinition struct {
	name string
	file string
	def  string
}

type CueRoot struct {
	dag.Root
}

func (r *CueRoot) PrintDag(i interface{}) {
	fmt.Print("Root:", r.Name, "\n\n")
	var ano func(f interface{}, indent int, m *dag.Node)

	ano = func(f interface{}, indent int, m *dag.Node) {
		if m.Value != nil {
			for i := 0; i < indent; i++ {
				fmt.Printf("%s", f)
			}
			v := m.Value.(NodeDefinition)
			fmt.Printf("%s is in %s\n", v.name, v.file)
		}
		for _, link := range m.Links {
			ano(f, indent+1, link)
		}
	}

	for _, m := range r.Members {
		ano(i, 0, m)
		fmt.Println("")
	}
}

func CreateCueDag(name string) CueRoot {
	return CueRoot{
		Root: dag.Root{Name: name},
	}
}
