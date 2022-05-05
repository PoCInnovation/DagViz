package main

import (
	"dag/lib"
)

func main() {
	/*root := lib.CreateDag("firstDag")
	node1 := lib.CreateNode(root, "TEST1")
	node2 := lib.CreateNode(root, []string{"info-1", "info-2"})
	lib.CreateNode(root, "LINK")
	node1.BothLinkTo(node2)
	root.PrintDag("---")*/

	root2 := lib.CreateDag("Alex")
	node3 := lib.CreateNode(root2, "Ismael")
	node4 := lib.CreateNode(root2, "Elie")
	node3.BothLinkTo(node4)
	root2.PrintDag("---")
}
