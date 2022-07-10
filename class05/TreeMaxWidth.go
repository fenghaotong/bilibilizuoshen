package class05

import (
	"algorithm/datastruct"
	"container/list"
	"fmt"
	"math"
)

func WidthTraversal(head *datastruct.BinaryTreeNode) {
	if head == nil {
		return
	}

	l := list.New()
	l.PushFront(head)

	for l.Len() != 0 {
		back := l.Back()
		node := back.Value.(*datastruct.BinaryTreeNode)
		fmt.Println(node.Val)
		if node.Left != nil {
			l.PushFront(node.Left)
		}

		if node.Right != nil {
			l.PushFront(node.Right)
		}

		if back != nil {
			l.Remove(back)
		}
	}

}

func TreeMaxWidthForMap(head *datastruct.BinaryTreeNode) int {
	if head == nil {
		return 0
	}

	l := list.New()
	l.PushFront(head)
	levelMap := make(map[*datastruct.BinaryTreeNode]int)
	levelMap[head] = 1
	curLevel := 1
	curLevelNodes := 0
	max := math.MinInt

	for l.Len() != 0 {
		back := l.Back()
		node := back.Value.(*datastruct.BinaryTreeNode)
		curNodeLevel := levelMap[node]
		if curNodeLevel == curLevel {
			curLevelNodes++
		} else {
			max = int(math.Max(float64(max), float64(curLevelNodes)))
			curLevel++
			curLevelNodes = 1
		}
		if node.Left != nil {
			levelMap[node.Left] = curNodeLevel + 1
			l.PushFront(node.Left)
		}

		if node.Right != nil {
			levelMap[node.Right] = curNodeLevel + 1
			l.PushFront(node.Right)
		}

		if back != nil {
			l.Remove(back)
		}
	}

	return max

}
