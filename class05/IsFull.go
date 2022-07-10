package class05

import (
	"algorithm/datastruct"
	"math"
)

// 满二叉树遍历
// nodes == 1 << height - 1

type CBTData struct {
	height int
	nodes  int
}

func IsCBTRecur(head *datastruct.BinaryTreeNode) bool {
	if head == nil {
		return true
	}
	cbt := processCBT(head)
	if cbt.nodes == (1<<cbt.height)-1 {
		return true
	}
	return false
}

func processCBT(head *datastruct.BinaryTreeNode) *CBTData {
	if head == nil {
		return &CBTData{0, 0}
	}

	leftData := processCBT(head.Left)
	rightData := processCBT(head.Right)

	height := int(math.Max(float64(leftData.height), float64(rightData.height))) + 1
	nodes := leftData.nodes + rightData.nodes + 1

	return &CBTData{height: height, nodes: nodes}
}
