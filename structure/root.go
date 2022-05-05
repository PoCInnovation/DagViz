package structure

import "fmt"

type Root struct {
	Name    string
	Members []*StructNode
}

func (r *Root) PrintDag(i interface{}) {
	fmt.Println("Root:", r.Name, "\n")
	for _, m := range r.Members {
		m.GetValue(i)
	}
}

func (r *Root) RemoveNode(node *StructNode) {

	for _, n := range r.Members {
		n.RemoveLink(node)
	}
	for i, n := range r.Members {
		if n == node {
			r.Members = append(r.Members[:i], r.Members[i+1:]...)
		}
	}
}
