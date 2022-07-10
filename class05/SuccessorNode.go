package class05

import "algorithm/datastruct"

// 后继节点

func GetSuccessNode(node *datastruct.BinaryTreeNodeS) *datastruct.BinaryTreeNodeS {
	if node == nil {
		return node
	}

	if node.Right != nil {
		return getMostLeft(node.Right)
	} else {
		parent := node.Parent
		if parent != nil && parent.Left != node {
			node = parent
			parent = node.Parent
		}
		return parent
	}
}

func getMostLeft(node *datastruct.BinaryTreeNodeS) *datastruct.BinaryTreeNodeS {
	if node == nil {
		return nil
	}

	for node.Left != nil {
		node = node.Left
	}
	return node
}
