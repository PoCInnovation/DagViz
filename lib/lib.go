package lib

import (
	s "dag/structure"
)

func CreateNode(root *s.Root, value interface{}) *s.StructNode {
	newNode := &s.StructNode{Value: value}
	root.Members = append(root.Members, newNode)
	return newNode
}

func CreateDag(name string) *s.Root {
	return &s.Root{Name: name}
}
