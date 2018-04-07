package tree

import (
	"reflect"
	"testing"
)

func TestTree(t *testing.T) {
	tree := NewEmptyTree().Insert(5).Insert(2).Insert(3)
	values := tree.Values()

	if !reflect.DeepEqual(values, []int{2, 3, 5}) {
		t.Error("on create tree should contain [2, 3, 5] but was", values)
	}

	if tree.GetNodeCount() != 3 {
		t.Error("on create expecting 3 values in the tree but was", tree.GetNodeCount())
	}

	if !tree.IsInTree(3) {
		t.Error("3 should be in tree")
	}

	if tree.IsInTree(10) {
		t.Error("10 should not be in tree")
	}

	if tree.GetMin() != 2 {
		t.Error("the minimum value should be 2 but was", tree.GetMin())
	}

	if tree.GetMax() != 5 {
		t.Error("the maximum value should be 5 but was", tree.GetMax())
	}

	if tree.GetHeight() != 3 {
		t.Error("tree height should be 3 but was", tree.GetHeight())
	}

	successor, ok := tree.GetSuccessor(5)
	if ok {
		t.Error("5 is the largest and shouldn't have successor but got", successor)
	}

	successor, ok = tree.GetSuccessor(3)
	if !ok {
		t.Error("3 should have a successor")
	}
	if successor != 5 {
		t.Error("5 is successor of 3 but got", successor)
	}

	successor, ok = tree.GetSuccessor(2)
	if !ok {
		t.Error("2 should have a successor")
	}
	if successor != 3 {
		t.Error("3 is successor of 2 but got", successor)
	}
}
