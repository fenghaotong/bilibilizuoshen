package class05

import "algorithm/datastruct"

// 树上两个节点的最低公共祖先

func LCARecur(head, o1, o2 *datastruct.BinaryTreeNode) *datastruct.BinaryTreeNode {

	if head == nil || head == o1 || head == o2 {
		return head
	}

	left := LCARecur(head.Left, o1, o2)
	right := LCARecur(head.Right, o1, o2)
	if left != nil && right != nil {
		return head
	}

	if left != nil {
		return left
	} else {
		return right
	}
}
