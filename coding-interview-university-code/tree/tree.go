package tree

import "math"

type BinarySearchTree struct {
	root *treeNode
}

type treeNode struct {
	Left  *treeNode
	Right *treeNode
	Value int
}

func NewEmptyTree() *BinarySearchTree {
	return &BinarySearchTree{}
}

func (tree *BinarySearchTree) Insert(value int) *BinarySearchTree {
	tree.root = insertRecursive(tree.root, value)
	return tree
}

func insertRecursive(node *treeNode, value int) *treeNode {
	if node == nil {
		return &treeNode{
			Left:  nil,
			Right: nil,
			Value: value,
		}
	}

	if value < node.Value {
		node.Left = insertRecursive(node.Left, value)
	} else if value > node.Value {
		node.Right = insertRecursive(node.Right, value)
	}

	return node
}

func (tree *BinarySearchTree) Values() []int {
	values := make([]int, 0)
	inOrderTraverse(tree.root, func(value int) bool {
		values = append(values, value)
		return true
	})
	return values
}

func (tree *BinarySearchTree) GetNodeCount() int {
	count := 0
	inOrderTraverse(tree.root, func(value int) bool {
		count++
		return true
	})
	return count
}

// TODO: do better than naive O(n)
func (tree *BinarySearchTree) GetSuccessor(value int) (int, bool) {
	captureNextValue := false

	successor := 0
	foundSuccessor := false

	inOrderTraverse(tree.root, func(currentValue int) bool {
		if captureNextValue {
			successor = currentValue
			foundSuccessor = true
			return false
		}
		if value == currentValue {
			captureNextValue = true
		}
		return true
	})

	return successor, foundSuccessor
}

func inOrderTraverse(node *treeNode, nodeVisitor func(int) bool) bool {
	keepTraversing := true

	if node != nil {
		keepTraversing = inOrderTraverse(node.Left, nodeVisitor)

		if keepTraversing {
			keepTraversing = nodeVisitor(node.Value)
		}

		if keepTraversing {
			keepTraversing = inOrderTraverse(node.Right, nodeVisitor)
		}
	}

	return keepTraversing
}

func (tree *BinarySearchTree) IsInTree(value int) bool {
	return findRecursive(tree.root, value)
}

func findRecursive(node *treeNode, value int) bool {
	if node == nil {
		return false
	}

	if value < node.Value {
		return findRecursive(node.Left, value)
	} else if value > node.Value {
		return findRecursive(node.Right, value)
	} else {
		return true
	}
}

func (tree *BinarySearchTree) GetMin() int {
	return getMinRecursive(tree.root)
}

func getMinRecursive(node *treeNode) int {
	if node.Left == nil {
		return node.Value
	}

	return getMinRecursive(node.Left)
}

func (tree *BinarySearchTree) GetMax() int {
	return getMaxRecursive(tree.root)
}

func getMaxRecursive(node *treeNode) int {
	if node.Right == nil {
		return node.Value
	}

	return getMaxRecursive(node.Right)
}

func (tree *BinarySearchTree) GetHeight() int {
	return int(getHeightRecursive(tree.root))
}

func getHeightRecursive(node *treeNode) float64 {
	if node == nil {
		return 0
	}

	heightLeft := 1 + getHeightRecursive(node.Left)
	heightRight := 1 + getHeightRecursive(node.Right)
	return math.Max(heightLeft, heightRight)
}
