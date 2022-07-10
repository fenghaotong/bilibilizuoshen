package class05

import (
	"algorithm/datastruct"
	"container/list"
)

// 完全二叉树判断
// 宽度优先遍历
// 1. 有右孩子没有左孩子，false
// 2. 有左孩子，左右孩子不是叶子节点，false

func IsCBT(head *datastruct.BinaryTreeNode) bool {
	if head == nil {
		return true
	}

	leaf := false
	queue := list.New()
	queue.PushFront(head)

	for queue.Len() != 0 {
		back := queue.Back()
		node := back.Value.(*datastruct.BinaryTreeNode)

		l := node.Left
		r := node.Right

		if (l == nil && r != nil) || (leaf && !(l == nil && r == nil)) {
			return false
		}
		if l != nil {
			queue.PushFront(l)
		}
		if r != nil {
			queue.PushFront(r)
		}

		if back != nil {
			queue.Remove(back)
		}

		if l == nil || r == nil {
			leaf = true
		}
	}
	return true
}
