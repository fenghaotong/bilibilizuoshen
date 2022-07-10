package class04

import "algorithm/datastruct"

func mergeTwoLists(list1 *datastruct.ListNode, list2 *datastruct.ListNode) *datastruct.ListNode {
	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	}
	head := &datastruct.ListNode{}
	cur := head

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			cur.Next = &datastruct.ListNode{Val: list1.Val}
			list1 = list1.Next
		} else {
			cur.Next = &datastruct.ListNode{Val: list2.Val}
			list2 = list2.Next
		}

		cur = cur.Next
	}

	if list1 != nil {
		cur.Next = list1

	}

	if list2 != nil {
		cur.Next = list2
	}

	return head.Next
}
