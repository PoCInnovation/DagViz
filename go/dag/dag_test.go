package dag

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDag(t *testing.T) {
	got := CreateDag("test")
	want := "test"

	assert.Equal(t, got.Name == want, true, "links still exist")
}

func TestGetValue(t *testing.T) {
	dag1 := CreateDag("test")
	testNode := dag1.AttachNode("testNode")
	want := "testNode"

	assert.Equalf(t, testNode.Value == want, true, "got %s but wanted %s", testNode.Value, "test")
}

func TestLinkTo(t *testing.T) {
	dag := CreateDag("test")
	testNode := dag.AttachNode("testNode")
	testNode2 := dag.AttachNode("testNode2")
	testNode3 := dag.AttachNode("testNode3")
	testNode.LinksTo(testNode2)
	testNode.LinksTo(testNode3)

	assert.Equalf(t, testNode.Links[0].Value == "testNode2", true, "got %s but wanted %s", testNode.Links[0].Value, "testNode2")
	assert.Equalf(t, testNode.Links[1].Value == "testNode3", true, "got %s but wanted %s", testNode.Links[1].Value, "testNode3")
}

func TestBothLinkTo(t *testing.T) {
	dag := CreateDag("test")
	testNode1 := dag.AttachNode("testNode1")
	testNode2 := dag.AttachNode("testNode2")
	testNode3 := dag.AttachNode("testNode3")
	testNode1.BothLinksTo(testNode2)
	testNode1.BothLinksTo(testNode3)

	assert.Equalf(t, testNode1.Links[0].Value == "testNode2", true, "got %s but wanted %s", testNode1.Links[0].Value, "testNode2")
	assert.Equalf(t, testNode2.Links[0].Value == "testNode1", true, "got %s but wanted %s", testNode2.Links[0].Value, "testNode1")
	assert.Equalf(t, testNode1.Links[1].Value == "testNode3", true, "got %s but wanted %s", testNode1.Links[1].Value, "testNode3")
	assert.Equalf(t, testNode3.Links[0].Value == "testNode1", true, "got %s but wanted %s", testNode3.Links[0].Value, "testNode1")
}

func TestRemoveLinkTo(t *testing.T) {
	dag := CreateDag("test")
	testNode := dag.AttachNode("testNode1")
	testNode2 := dag.AttachNode("testNode2")
	testNode3 := dag.AttachNode("testNode3")
	testNode.BothLinksTo(testNode2)
	testNode.BothLinksTo(testNode3)
	testNode.RemoveLinkTo(testNode2)
	testNode.RemoveLinkTo(testNode3)
	testNode2.RemoveLinkTo(testNode)

	assert.Equal(t, len(testNode.Links) == 0, true, "links still exist")
	assert.Equal(t, len(testNode2.Links) == 0, true, "links still exist")
	assert.Equal(t, len(testNode3.Links) != 0, true, "links do not exist")
}

func TestRemoveNode(t *testing.T) {
	dag := CreateDag("test")
	testNode := dag.AttachNode("testNode1")
	testNode2 := dag.AttachNode("testNode2")
	testNode3 := dag.AttachNode("testNode3")
	testNode.BothLinksTo(testNode2)
	testNode.BothLinksTo(testNode3)
	dag.RemoveNode(testNode)

	assert.Equal(t, len(testNode2.Links) == 0, true, "links still exist")
	assert.Equal(t, len(testNode3.Links) == 0, true, "links still exist")
}
