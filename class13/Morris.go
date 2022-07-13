package class13

import (
	"algorithm/datastruct"
	"fmt"
)

func Morris(head *datastruct.BinaryTreeNode) {
	if head == nil {
		return
	}

	cur := head
	mostRight := new(datastruct.BinaryTreeNode)
	for cur != nil {

		mostRight = cur.Left
		if mostRight == nil {
			cur = cur.Right
		} else {
			for mostRight.Right != nil && mostRight.Right != cur {
				mostRight = mostRight.Right
			}

			if mostRight.Right == nil {
				mostRight.Right = cur
				cur = cur.Left

			} else {
				mostRight.Right = nil
				cur = cur.Right
			}
		}
	}
}

func PreMorris(head *datastruct.BinaryTreeNode) {
	if head == nil {
		return
	}

	cur := head
	mostRight := new(datastruct.BinaryTreeNode)
	for cur != nil {
		mostRight = cur.Left
		if mostRight != nil {
			for mostRight.Right != nil && mostRight.Right != cur {
				mostRight = mostRight.Right
			}

			if mostRight.Right == nil {
				fmt.Println(cur.Val)
				mostRight.Right = cur
				cur = cur.Left
				continue
			} else {
				mostRight.Right = nil
			}
		} else {
			fmt.Println(cur.Val)
		}
		cur = cur.Right
	}
}

func InMorris(head *datastruct.BinaryTreeNode) {
	if head == nil {
		return
	}

	cur := head
	mostRight := new(datastruct.BinaryTreeNode)
	for cur != nil {
		mostRight = cur.Left
		if mostRight != nil {
			for mostRight.Right != nil && mostRight.Right != cur {
				mostRight = mostRight.Right
			}

			if mostRight.Right == nil {
				mostRight.Right = cur
				cur = cur.Left
				continue
			} else {
				mostRight.Right = nil
			}
		}

		fmt.Println(cur.Val)
		cur = cur.Right
	}
}

func PostMorris(head *datastruct.BinaryTreeNode) {
	if head == nil {
		return
	}

	cur := head
	mostRight := new(datastruct.BinaryTreeNode)
	for cur != nil {
		mostRight = cur.Left
		if mostRight != nil {
			for mostRight.Right != nil && mostRight.Right != cur {
				mostRight = mostRight.Right
			}

			if mostRight.Right == nil {
				mostRight.Right = cur
				cur = cur.Left
				continue
			} else {
				mostRight.Right = nil
				printRight(cur.Left)
			}
		}
		cur = cur.Right
	}
	printRight(head)
}

func printRight(node *datastruct.BinaryTreeNode) {
	tail := reverseNode(node)
	cur := tail
	for cur != nil {
		if cur.Val != 0 {
			fmt.Print(cur.Val)
			fmt.Print(" ")
		}

		cur = cur.Right
	}
	reverseNode(tail)
	fmt.Print(" ")
}

func reverseNode(node *datastruct.BinaryTreeNode) *datastruct.BinaryTreeNode {
	pre := new(datastruct.BinaryTreeNode)
	next := new(datastruct.BinaryTreeNode)
	for node != nil {
		next = node.Right
		node.Right = pre
		pre = node
		node = next
	}
	return pre
}
