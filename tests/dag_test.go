package tests

import (
	"dagviz/dag"
	"testing"
)

func TestDag(t *testing.T) {
	got := dag.CreateDag("test")
	want := "test"

	if got.Name != want {
		t.Errorf("got %s but wanted %s", got.Name, want)
	}
}

func TestGetValue(t *testing.T) {
	dag1 := dag.CreateDag("test")
	testNode := dag1.AttachNode("testNode")

	if testNode.Value != "testNode" {
		t.Errorf("got %s but wanted %s", testNode.Value, "test")
	}
}

func TestLinkTo(t *testing.T) {
	dag := dag.CreateDag("test")
	testNode := dag.AttachNode("testNode")
	testNode2 := dag.AttachNode("testNode2")
	testNode3 := dag.AttachNode("testNode3")
	testNode.LinksTo(testNode2)
	testNode.LinksTo(testNode3)

	if testNode.Links[0].Value != "testNode2" {
		t.Errorf("got %s but wanted %s", testNode.Links[0].Value, "testNode2")
	}
	if testNode.Links[1].Value != "testNode3" {
		t.Errorf("got %s but wanted %s", testNode.Links[1].Value, "testNode3")
	}
}

func TestBothLinkTo(t *testing.T) {
	dag := dag.CreateDag("test")
	testNode1 := dag.AttachNode("testNode1")
	testNode2 := dag.AttachNode("testNode2")
	testNode3 := dag.AttachNode("testNode3")
	testNode1.BothLinksTo(testNode2)
	testNode1.BothLinksTo(testNode3)

	if testNode1.Links[0].Value != "testNode2" {
		t.Errorf("got %s but wanted %s", testNode1.Links[0].Value, "testNode2")
	}
	if testNode2.Links[0].Value != "testNode1" {
		t.Errorf("got %s but wanted %s", testNode2.Links[0].Value, "testNode1")
	}
	if testNode1.Links[1].Value != "testNode3" {
		t.Errorf("got %s but wanted %s", testNode1.Links[1].Value, "testNode3")
	}
	if testNode3.Links[0].Value != "testNode1" {
		t.Errorf("got %s but wanted %s", testNode3.Links[0].Value, "testNode1")
	}
}

func TestRemoveLinkTo(t *testing.T) {
	dag := dag.CreateDag("test")
	testNode := dag.AttachNode("testNode1")
	testNode2 := dag.AttachNode("testNode2")
	testNode3 := dag.AttachNode("testNode3")
	testNode.BothLinksTo(testNode2)
	testNode.BothLinksTo(testNode3)
	testNode.RemoveLinkTo(testNode2)
	testNode.RemoveLinkTo(testNode3)
	testNode2.RemoveLinkTo(testNode)

	if len(testNode.Links) > 0 {
		t.Errorf("links still exist")
	}
	if len(testNode2.Links) > 0 {
		t.Errorf("links still exist")
	}
	if len(testNode3.Links) < 0 {
		t.Errorf("links do not exist")
	}
}

func TestRemoveNode(t *testing.T) {
	dag := dag.CreateDag("test")
	testNode := dag.AttachNode("testNode1")
	testNode2 := dag.AttachNode("testNode2")
	testNode3 := dag.AttachNode("testNode3")
	testNode.BothLinksTo(testNode2)
	testNode.BothLinksTo(testNode3)
	dag.RemoveNode(testNode)

	if len(testNode2.Links) > 0 {
		t.Errorf("links still exist")
	}
	if len(testNode3.Links) > 0 {
		t.Errorf("links still exist")
	}
}
