package main

import (
	"dag/lib"
)

func main() {
	/*root := lib.CreateDag("firstDag")
	node1 := lib.AttachNode(root, "TEST1")
	node2 := lib.AttachNode(root, []string{"info-1", "info-2"})
	lib.AttachNode(root, "LINK")
	node1.BothLinksTo(node2)
	root.PrintDag("---")*/

	root2 := lib.CreateDag("Alex")
	node3 := root2.AttachNode("Ismael")
	node4 := root2.AttachNode("Elie")
	node3.BothLinksTo(node4)
	root2.PrintDag("---")
}
