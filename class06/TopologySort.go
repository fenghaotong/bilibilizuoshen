package class06

import (
	"algorithm/datastruct"
	"container/list"
)

// 拓扑排序

func TopopogySort(graph datastruct.Graph) *list.List {
	inMap := make(map[datastruct.GraphNode]int, 0)
	zeroInQueue := list.New()

	for _, v := range graph.Nodes {
		inMap[v] = v.In
		if v.In == 0 { //  入度为1的节点
			zeroInQueue.PushFront(v)
		}
	}

	list := list.New()
	for zeroInQueue.Len() != 0 {
		back := zeroInQueue.Back()
		node := back.Value.(datastruct.GraphNode)
		list.PushBack(node)
		for _, next := range *node.Next {
			inMap[next] = inMap[next] - 1
			if inMap[next] == 0 {
				zeroInQueue.PushFront(next)
			}
		}

		if back != nil {
			zeroInQueue.Remove(back)
		}
	}
	return list
}
