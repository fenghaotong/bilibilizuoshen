package class04

import (
	"algorithm/datastruct"
)

// 将一个链表按照某个值大于放左边，小于放右边， 等于放中间

func SwapObj(arr []datastruct.ListNode, index1, index2 int) {
	tmp := arr[index1]
	arr[index1] = arr[index2]
	arr[index2] = tmp
}

// ListPartition1
// 使用一个数组将所有的节点放入到数组中，使用快排进行排序
func ListPartition1(list1 *datastruct.LinkList, pivot datastruct.Element) *datastruct.LinkList {
	if list1.Head == nil || list1.Head.Next == nil {
		return list1
	}
	helps := make([]datastruct.ListNode, 0) // 辅助空间
	head := list1.Head
	for head.Next != nil {
		head = head.Next
		helps = append(helps, *head)
	}

	SwapObj(helps, len(helps)-1, GetIndex(helps, pivot))
	Process(helps, 0, len(helps)-1)

	list := datastruct.NewLinkList()
	for _, value := range helps {
		list.AddBack(value.Val)
	}
	return list
}

func GetIndex(arr []datastruct.ListNode, node datastruct.Element) int {
	cur := datastruct.ListNode{Val: node}
	for index, value := range arr {
		if value.Equal(cur) {
			return index
		}
	}
	return len(arr) - 1
}

func Process(arr []datastruct.ListNode, L, R int) {
	if L < R {
		p := Partition(arr, L, R)
		Process(arr, L, p[0]-1)
		Process(arr, p[1]+1, R)
	}
}

func Partition(arr []datastruct.ListNode, L, R int) []int {

	less := L - 1
	more := R
	for L < more {
		if arr[L].Less(arr[R]) {
			less++
			SwapObj(arr, less, L)
			L++
		} else if arr[L].More(arr[R]) {
			more--
			SwapObj(arr, more, L)
		} else {
			L++
		}
	}

	SwapObj(arr, more, R)
	return []int{less + 1, more}
}

// ListPartition2
// 使用有限空间实现
func ListPartition2(list1 *datastruct.LinkList, pivot datastruct.Element) *datastruct.LinkList {
	lessH := new(datastruct.ListNode)
	lessT := new(datastruct.ListNode)
	equalH := new(datastruct.ListNode)
	equalT := new(datastruct.ListNode)
	moreH := new(datastruct.ListNode)
	moreT := new(datastruct.ListNode)
	lessH = nil
	lessT = nil
	equalH = nil
	equalT = nil
	moreH = nil
	moreT = nil

	head := list1.Head
	cur := datastruct.ListNode{Val: pivot}
	for head != nil {
		next := head.Next
		head.Next = nil
		if head.Less(cur) {
			if lessH == nil {
				lessH = head
				lessT = head
				lessH.Next = lessT
			} else {
				lessT.Next = head
				lessT = head
			}
		}

		if head.Equal(cur) {
			if equalH == nil {
				equalH = head
				equalT = head
				equalH.Next = equalT
			} else {
				equalT.Next = head
				equalT = head
			}

		}

		if head.More(cur) {
			if moreH == nil {
				moreH = head
				moreT = head
				moreH.Next = moreT
			} else {
				moreT.Next = head
				moreT = head
			}
		}
		head = next
	}

	if lessT != nil {
		lessT.Next = equalH
		if equalT == nil {
			equalT = lessT
		}
	}

	if equalT != nil {
		equalT.Next = moreH
	}

	list := datastruct.NewLinkList()
	if lessH != nil {
		list.Head = lessH
	} else if equalH != nil {
		list.Head = equalH
	} else {
		list.Head = moreH
	}

	return list

}
