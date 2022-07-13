package class13

import (
	"algorithm/datastruct"
	"testing"
)

func InitTree1() *datastruct.BinaryTreeNode {
	head := &datastruct.BinaryTreeNode{Val: 1}
	head.Left = &datastruct.BinaryTreeNode{Val: 2}
	head.Right = &datastruct.BinaryTreeNode{Val: 3}
	head.Left.Left = &datastruct.BinaryTreeNode{Val: 4}
	head.Left.Right = &datastruct.BinaryTreeNode{Val: 5}
	head.Right.Left = &datastruct.BinaryTreeNode{Val: 6}
	head.Right.Right = &datastruct.BinaryTreeNode{Val: 7}
	head.Left.Left.Left = &datastruct.BinaryTreeNode{Val: 8}
	head.Left.Left.Right = &datastruct.BinaryTreeNode{Val: 9}
	head.Left.Right.Left = &datastruct.BinaryTreeNode{Val: 10}
	//head.Left.Right.Right = &datastruct.BinaryTreeNode{Val: 11}
	datastruct.PrintTree(head)
	return head
}

func TestMorris(t *testing.T) {
	tree1 := InitTree1()

	Morris(tree1)
}

func TestPreMorris(t *testing.T) {
	tree1 := InitTree1()

	PreMorris(tree1)
}

func TestInMorris(t *testing.T) {
	tree1 := InitTree1()

	InMorris(tree1)
}

func TestPostMorris(t *testing.T) {
	tree1 := InitTree1()

	PostMorris(tree1)
}
