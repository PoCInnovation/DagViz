package cueToDag

import (
	"dagviz/dag"
	"fmt"
	"strings"
)

func CreateCueDag(name string) CueRoot {
	return CueRoot{
		Root: dag.Root{Name: name},
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

func (r *CueRoot) PrintDag(i interface{}) {
	fmt.Print("Root:", r.Name, "\n\n")

	for _, m := range r.Members {
		printNode(i, 0, m)
		fmt.Println("")
	}
}

func formatDef(s string) string {
	s = strings.Replace(s, "\"", "\\\"", -1)
	s = strings.Replace(s, "\n", "\\n", -1)
	s = strings.Replace(s, "\t", "\\t", -1)
	return s
}

func fmtPrintCueJson(d NodeDefinition, links []*dag.Node) {
	fmt.Printf("{\"name\":\"%s\", \"file\":\"%s\", \"def\":\"%s\", \"dependencies\": ", d.name, d.file, formatDef(d.def))
	PrintJson(links)
	fmt.Print("}")
}

func PrintJson(members []*dag.Node) {
	fmt.Print("[")

	for index, m := range members {
		if m.Value != nil {
			if index != 0 {
				fmt.Print(",")
			}
			v := m.Value.(NodeDefinition)
			fmtPrintCueJson(v, m.Links)
		}

	}
	fmt.Print("]")
}
