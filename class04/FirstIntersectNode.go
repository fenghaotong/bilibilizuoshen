package class04

import (
	"algorithm/datastruct"
	"math"
)

func GetIntersectNode(list1 *datastruct.LinkList, list2 *datastruct.LinkList) *datastruct.ListNode {
	if list1.Head == nil || list2.Head == nil {
		return nil
	}

	loop1 := GetLoopNode(list1.Head)
	loop2 := GetLoopNode(list2.Head)

	if loop1 == nil && loop2 == nil {
		return NoLoop(list1.Head, list2.Head)
	}

	if loop1 != nil && loop2 != nil {
		return BothLoop(list1.Head, loop1, list2.Head, loop2)
	}

	return nil
}

func GetLoopNode(head *datastruct.ListNode) *datastruct.ListNode {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return nil
	}

	s := head.Next
	f := head.Next.Next
	for s != f {
		if f.Next == nil || f.Next.Next == nil {
			return nil
		}
		s = s.Next
		f = f.Next.Next
	}

	f = head
	for s != f {
		s = s.Next
		f = f.Next
	}
	return s
}

func NoLoop(head1 *datastruct.ListNode, head2 *datastruct.ListNode) *datastruct.ListNode {
	count := 0
	cur1 := head1
	cur2 := head2

	for cur1 != nil {
		count++
		cur1 = cur1.Next
	}

	for cur2 != nil {
		count--
		cur2 = cur2.Next
	}

	if cur1 != cur2 {
		return nil
	}

	if count > 0 {
		cur1 = head1
		cur2 = head2
	} else {
		cur1 = head2
		cur2 = head1
	}

	count = int(math.Abs(float64(count)))
	for count != 0 {
		cur1 = cur1.Next
		count--
	}

	for cur1 != cur2 {
		cur1 = cur1.Next
		cur2 = cur2.Next
	}

	return cur1
}

func BothLoop(head1 *datastruct.ListNode, loop1 *datastruct.ListNode, head2 *datastruct.ListNode, loop2 *datastruct.ListNode) *datastruct.ListNode {
	if loop1 == loop2 {
		count := 0
		cur1 := head1
		cur2 := head2

		for cur1 != loop1 {
			count++
			cur1 = cur1.Next
		}

		for cur2 != loop1 {
			count--
			cur2 = cur2.Next
		}

		if count > 0 {
			cur1 = head1
			cur2 = head2
		} else {
			cur1 = head2
			cur2 = head1
		}

		count = int(math.Abs(float64(count)))
		for count != 0 {
			cur1 = cur1.Next
			count--
		}

		for cur1 != cur2 {
			cur1 = cur1.Next
			cur2 = cur2.Next
		}
		return cur1
	} else {
		cur1 := loop1.Next
		for cur1 != loop1 {
			if cur1 == loop2 {
				return loop1
			}
			cur1 = cur1.Next
		}
		return nil
	}
}
