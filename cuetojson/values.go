package cuetojson

import (
	"cuelang.org/go/cue"
	"dagviz/dag"
)

func AppendValuesToDag(root *dag.Root, programs []CueProgram) {
	for _, program := range programs {
		addNode(root, nil, program.Values, nil)
	}
	root.PrintDag(" --> ")
}

func addNode(root *dag.Root, node *dag.Node, value *cue.Value, prev *cue.Iterator) {
	iterator, err := value.Fields()

	if err != nil {
		if iterator != nil {
			addDagEdge(root, node, prev.Label(), value)
		}
	}

	for iterator.Next() {
		v := iterator.Value()
		_, err = v.Fields()

		if err != nil {
			addNode(root, node, &v, iterator)
		} else {
			addNode(root, addToDag(root, node, iterator.Label()), &v, iterator)
		}
	}
}

func addToDag(root *dag.Root, node *dag.Node, content string) *dag.Node {
	if node == nil {
		return root.AttachNode(content)
	}

	newNode := &dag.Node{Value: content}
	node.LinksTo(newNode)
	return newNode
}

func addDagEdge(root *dag.Root, node *dag.Node, key string, value *cue.Value) {
	v, err := value.String()
	if err != nil {
		return
	}
	format := key + " = " + v
	addToDag(root, node, format)
}
