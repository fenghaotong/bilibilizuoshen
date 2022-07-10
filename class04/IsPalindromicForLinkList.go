package class04

import (
	"algorithm/datastruct"
)

// IsPalindromic1
// 使用栈将所有数据放入栈中
func IsPalindromic1(list1 *datastruct.LinkList) bool {
	if list1.Head == nil || list1.Head.Next == nil {
		return true
	}
	newList := datastruct.NewLinkList()
	node := list1.Copy()
	l := node.Head
	for l.Next != nil {
		l = l.Next
		newList.AddFront(l.Val)
	}

	l = node.Head.Next
	for newList.Head.Next != nil {
		front := newList.RemoveFront()
		if front.Val != l.Val {
			return false
		}
		l = l.Next
	}
	return true
}

// IsPalindromic2
// 使用栈将一半数据放入栈中，使用快慢指针辅助
func IsPalindromic2(list1 *datastruct.LinkList) bool {
	if list1.Head == nil || list1.Head.Next == nil {
		return true
	}
	newList := datastruct.NewLinkList()
	node := list1.Copy()
	s := node.Head.Next
	f := node.Head
	for f.Next != nil && f.Next.Next != nil {
		s = s.Next
		f = f.Next.Next
	}

	for s != nil {
		newList.AddFront(s.Val)
		s = s.Next
	}
	l := node.Head.Next
	for newList.Head.Next != nil {
		front := newList.RemoveFront()
		if front.Val != l.Val {
			return false
		}
		l = l.Next
	}
	return true
}

// IsPalindromic3
// 使用快慢指针，不适用额外的空间实现
func IsPalindromic3(list1 *datastruct.LinkList) bool {
	if list1.Head == nil || list1.Head.Next == nil {
		return true
	}
	node := list1.Copy()
	s := node.Head
	f := node.Head
	for f.Next != nil && f.Next.Next != nil {
		s = s.Next
		f = f.Next.Next
	}

	f = s.Next
	s.Next = nil
	tmp := &datastruct.ListNode{}
	for f != nil {
		tmp = f.Next
		f.Next = s
		s = f
		f = tmp
	}

	tmp = s
	f = node.Head.Next
	res := true
	for {
		if s != nil && f != nil {
			if s.Val != f.Val {
				res = false
				break
			}
			s = s.Next
			f = f.Next
		} else {
			break
		}
	}

	s = tmp.Next
	tmp.Next = nil
	for s != nil {
		f = s.Next
		s.Next = tmp
		tmp = s
		s = f
	}

	return res
}
