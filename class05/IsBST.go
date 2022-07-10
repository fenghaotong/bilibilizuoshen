package class05

import (
	"algorithm/datastruct"
	"math"
)

// 搜索二叉树判断
// 二叉树中序遍历是否有序

var preValue datastruct.Element = math.MinInt

func CheckBstRecur(head *datastruct.BinaryTreeNode) bool {

	if head == nil {
		return true
	}

	isLeftBst := CheckBstRecur(head.Left)
	if !isLeftBst {
		return false
	}

	if head.Val <= preValue {
		return false
	} else {
		preValue = head.Val
	}

	return CheckBstRecur(head.Right)

}

func CheckBst(head *datastruct.BinaryTreeNode) bool {
	if head == nil {
		return true
	}

	stack := datastruct.NewStack()
	var tmp datastruct.Element = math.MinInt

	for !stack.Empty() || head != nil {
		if head != nil {
			stack.Push(head)
			head = head.Left
		} else {
			head = stack.Pop().(*datastruct.BinaryTreeNode)
			if head.Val <= tmp {
				return false
			} else {
				tmp = head.Val
			}
			head = head.Right
		}
	}
	return true
}

type BSTData struct {
	isBst bool
	min   datastruct.Element
	max   datastruct.Element
}

func IsBst(head *datastruct.BinaryTreeNode) bool {
	return processBST(head).isBst
}

func processBST(head *datastruct.BinaryTreeNode) *BSTData {
	if head == nil {
		return nil
	}

	leftData := processBST(head.Left)
	rightData := processBST(head.Right)

	isBst := true
	min := head.Val
	max := head.Val

	if leftData != nil {
		min = datastruct.Element(math.Min(float64(min), float64(leftData.min)))
		max = datastruct.Element(math.Max(float64(max), float64(leftData.max)))
	}

	if rightData != nil {
		min = datastruct.Element(math.Min(float64(min), float64(rightData.min)))
		max = datastruct.Element(math.Max(float64(max), float64(rightData.max)))
	}

	if leftData != nil && (!leftData.isBst || leftData.max >= head.Val) {
		isBst = false
	}

	if rightData != nil && (!rightData.isBst || rightData.min <= head.Val) {
		isBst = false
	}
	return &BSTData{isBst: isBst, min: min, max: max}
}
