package class05

import (
	"algorithm/datastruct"
	"math"
)

// 平衡二叉树 AVL

type ReturnType struct {
	height    int
	isBalance bool
}

func IsBalancedTree(head *datastruct.BinaryTreeNode) bool {
	return processBT(head).isBalance
}

func processBT(head *datastruct.BinaryTreeNode) ReturnType {
	if head == nil {
		return ReturnType{0, true}
	}
	leftData := processBT(head.Left)
	rightData := processBT(head.Right)

	height := int(math.Max(float64(leftData.height), float64(rightData.height))) + 1
	isBalance := false
	if leftData.isBalance && rightData.isBalance && math.Abs(float64(rightData.height-leftData.height)) < 2 {
		isBalance = true
	}

	return ReturnType{height: height, isBalance: isBalance}
}
