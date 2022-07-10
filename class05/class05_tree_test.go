package class05

import (
	"algorithm/datastruct"
	"fmt"
	"testing"
)

/*
					1
			2				3
		4		5		6    	7
	8     9  10  11
*/

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

func InitBST() *datastruct.BinaryTreeNode {
	head := &datastruct.BinaryTreeNode{Val: 5}
	head.Left = &datastruct.BinaryTreeNode{Val: 3}
	head.Right = &datastruct.BinaryTreeNode{Val: 8}
	head.Left.Left = &datastruct.BinaryTreeNode{Val: 2}
	head.Left.Right = &datastruct.BinaryTreeNode{Val: 4}
	head.Right.Left = &datastruct.BinaryTreeNode{Val: 7}
	head.Right.Right = &datastruct.BinaryTreeNode{Val: 9}
	datastruct.PrintTree(head)
	return head
}

func InitTreeS() *datastruct.BinaryTreeNodeS {
	head := &datastruct.BinaryTreeNodeS{Val: 1}
	head.Left = &datastruct.BinaryTreeNodeS{Val: 2}
	head.Left.Parent = head
	head.Right = &datastruct.BinaryTreeNodeS{Val: 3}
	head.Right.Parent = head
	head.Left.Left = &datastruct.BinaryTreeNodeS{Val: 4}
	head.Left.Left.Parent = head.Left
	head.Left.Right = &datastruct.BinaryTreeNodeS{Val: 5}
	head.Left.Right.Parent = head.Left
	head.Right.Left = &datastruct.BinaryTreeNodeS{Val: 6}
	head.Right.Left.Parent = head.Parent
	head.Right.Right = &datastruct.BinaryTreeNodeS{Val: 7}
	head.Right.Right.Parent = head.Parent
	//datastruct.PrintTree(head)
	return head
}

func TestPreOrderRecur(t *testing.T) {
	head := InitTree1()
	PreOrderRecur(head)
}

func TestPreOrder(t *testing.T) {
	head := InitTree1()
	PreOrder(head)
}

func TestInOrderRecur(t *testing.T) {
	head := InitTree1()
	InOrderRecur(head)
}

func TestInOrder(t *testing.T) {
	head := InitTree1()
	InOrder(head)
}

func TestPosOrderRecur(t *testing.T) {
	head := InitTree1()
	PosOrderRecur(head)
}

func TestPosOrder(t *testing.T) {
	head := InitTree1()
	PosOrder(head)
}

func TestWidthTraversal(t *testing.T) {
	head := InitTree1()
	WidthTraversal(head)
}

func TestTreeMaxWidth(t *testing.T) {
	head := InitTree1()
	fmt.Println(TreeMaxWidthForMap(head))
}

func TestCheckBst(t *testing.T) {
	head := InitBST()
	fmt.Println(CheckBst(head))
	fmt.Println(CheckBstRecur(head))
	fmt.Println(IsBst(head))
}

func TestIsCBT(t *testing.T) {
	head := InitTree1()
	fmt.Println(IsCBT(head))
	fmt.Println(IsCBTRecur(head))
}

func TestIsBalancedTree(t *testing.T) {
	head := InitTree1()
	fmt.Println(IsBalancedTree(head))
}

func TestLCARecur(t *testing.T) {
	head := InitTree1()
	fmt.Println(LCARecur(head, head.Right.Left, head.Left.Right))
}

func TestGetSuccessNode(t *testing.T) {
	head := InitTreeS()
	fmt.Println(GetSuccessNode(head.Left))
}

func TestPrintPaperFolding(t *testing.T) {
	PrintPaperFolding(3)
}
