package class05

import (
	"algorithm/datastruct"
	"fmt"
)

func PreOrderRecur(head *datastruct.BinaryTreeNode) {
	if head == nil {
		return
	}
	fmt.Println(head.Val)
	PreOrderRecur(head.Left)
	PreOrderRecur(head.Right)
}

func PreOrder(head *datastruct.BinaryTreeNode) {
	if head == nil {
		return
	}

	stack := datastruct.NewStack()

	stack.Push(head)

	for !stack.Empty() {
		node := stack.Pop().(*datastruct.BinaryTreeNode)
		fmt.Println(node.Val)
		if node.Right != nil {
			stack.Push(node.Right)
		}

		if node.Right != nil {
			stack.Push(node.Left)
		}
	}
}

func InOrderRecur(head *datastruct.BinaryTreeNode) {
	if head == nil {
		return
	}

	InOrderRecur(head.Left)
	fmt.Println(head.Val)
	InOrderRecur(head.Right)
}

func InOrder(head *datastruct.BinaryTreeNode) {
	if head == nil {
		return
	}
	stack := datastruct.NewStack()
	for !stack.Empty() || head != nil {
		if head != nil {
			stack.Push(head)
			head = head.Left
		} else {
			head = stack.Pop().(*datastruct.BinaryTreeNode)
			fmt.Println(head.Val)
			head = head.Right
		}
	}
}

func PosOrderRecur(head *datastruct.BinaryTreeNode) {
	if head == nil {
		return
	}

	PosOrderRecur(head.Left)
	PosOrderRecur(head.Right)
	fmt.Println(head.Val)
}

func PosOrder(head *datastruct.BinaryTreeNode) {
	if head == nil {
		return
	}

	stack := datastruct.NewStack()
	helpStack := datastruct.NewStack()
	stack.Push(head)

	for !stack.Empty() {
		node := stack.Pop().(*datastruct.BinaryTreeNode)
		helpStack.Push(node)
		if node.Left != nil {
			stack.Push(node.Left)
		}

		if node.Right != nil {
			stack.Push(node.Right)
		}
	}

	for !helpStack.Empty() {
		helpNode := helpStack.Pop().(*datastruct.BinaryTreeNode)
		fmt.Println(helpNode.Val)
	}
}
