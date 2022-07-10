package datastruct

import (
	"fmt"
	"strings"
)

type BinaryTreeNode struct {
	Val   Element
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

// 特殊的树

type BinaryTreeNodeS struct {
	Val    Element
	Left   *BinaryTreeNodeS
	Right  *BinaryTreeNodeS
	Parent *BinaryTreeNodeS
}

func PrintTree(head *BinaryTreeNode) {
	fmt.Println("Binary Tree:")
	printInOrder(head, 0, "H", 17)
	fmt.Println()
}

func printInOrder(head *BinaryTreeNode, height int, to string, length int) {
	if head == nil {
		return
	}
	printInOrder(head.Right, height+1, "v", length)
	val := fmt.Sprintf("%s%v%s", to, head.Val, to)
	lenM := len(val)
	lenL := (length - lenM) / 2
	lenR := length - lenM - lenL
	val = getSpace(lenL) + val + getSpace(lenR)
	fmt.Println(getSpace(height*length) + val)
	printInOrder(head.Left, height+1, "^", length)
}

func getSpace(num int) string {
	space := " "
	buf := make([]string, 0)
	for i := 0; i < num; i++ {
		buf = append(buf, space)
	}
	return strings.Join(buf, "")
}
