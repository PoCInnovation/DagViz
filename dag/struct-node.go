package dag

import "fmt"

type Node struct {
	Value interface{}
	Links []*Node
}

func (n *Node) GetValue(i interface{}) {
	if n.Value != nil {
		fmt.Println(i, n.Value)
	}

	for _, link := range n.Links {
		if link.Value != nil {
			fmt.Println(i, "--->", link.Value, "\n")
		}
	}
}

func (n *Node) LinksTo(node ...*Node) {
	for _, l := range node {
		n.Links = append(n.Links, l)
	}
}

func (n *Node) BothLinksTo(node ...*Node) {
	for _, l := range node {
		n.Links = append(n.Links, l)
		l.Links = append(l.Links, n)
	}
}

func (n *Node) RemoveLinkTo(node *Node) {
	for i, l := range n.Links {
		if l == node {
			n.Links = append(n.Links[:i], n.Links[i+1:]...)
			break
		}
	}
}
