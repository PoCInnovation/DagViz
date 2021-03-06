package dag

import (
	"fmt"
)

type Root struct {
	Name    string
	Members []*Node
}

func (r *Root) PrintDag(i interface{}) {
	fmt.Print("Root:", r.Name, "\n\n")
	for _, m := range r.Members {
		m.GetValue(i, 0)
		fmt.Println("")
	}
}

func (r *Root) RemoveNode(node *Node) {
	for _, n := range r.Members {
		n.RemoveLinkTo(node)
	}

	for i, n := range r.Members {
		if n == node {
			r.Members = append(r.Members[:i], r.Members[i+1:]...)
		}
	}
}

func (r *Root) AttachNode(value interface{}) *Node {
	newNode := &Node{Value: value}
	r.Members = append(r.Members, newNode)
	return newNode
}
