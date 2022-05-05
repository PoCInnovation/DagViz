package structure

import "fmt"

type StructNode struct {
	Value interface{}
	Links []*StructNode
}

func (n *StructNode) GetValue(i interface{}) {
	if n.Value != nil {
		fmt.Println(i, n.Value)
	}
	for _, link := range n.Links {
		if link.Value != nil {
			fmt.Println(i, "--->", link.Value, "\n")
		}
	}
}

func (n *StructNode) LinkTo(node ...*StructNode) {
	for _, l := range node {
		n.Links = append(n.Links, l)
	}
}

func (n *StructNode) BothLinkTo(node ...*StructNode) {
	for _, l := range node {
		n.Links = append(n.Links, l)
		l.Links = append(l.Links, n)
	}
}

func (n *StructNode) RemoveLink(node *StructNode) {
	for i, l := range n.Links {
		if l == node {
			n.Links = append(n.Links[:i], n.Links[i+1:]...)
			break
		}
	}
}
