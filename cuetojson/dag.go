package cuetojson

import (
	"dagviz/dag"
	"fmt"
)

func CreateCueDag(name string) CueRoot {
	return CueRoot{
		Root: dag.Root{Name: name},
	}
}

func (r *CueRoot) PrintDag(i interface{}) {
	fmt.Print("Root:", r.Name, "\n\n")

	for _, m := range r.Members {
		printNode(i, 0, m)
		fmt.Println("")
	}
}

func printNode(f interface{}, indent int, m *dag.Node) {
	if m.Value != nil {
		for i := 0; i < indent; i++ {
			fmt.Printf("%s", f)
		}
		v := m.Value.(NodeDefinition)
		fmt.Printf("%s is in %s\n", v.name, v.file)
	}
	for _, link := range m.Links {
		printNode(f, indent+1, link)
	}
}
