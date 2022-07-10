package datastruct

import "fmt"

type Element int

type ListNode struct {
	Val  Element
	Next *ListNode
}

type LinkList struct {
	Head *ListNode
}

func NewLinkList() *LinkList {
	return new(LinkList).Init()
}

func (l *LinkList) Init() *LinkList {
	l.Head = &ListNode{}
	return l
}

func (l *LinkList) AddBack(data Element) {
	head := l.Head
	for {
		if head.Next == nil {
			node := &ListNode{Val: data}
			head.Next = node
			break
		} else {
			head = head.Next
		}
	}
}

func (l *LinkList) AddBackElements(data []Element) {
	for _, v := range data {
		l.AddBack(v)
	}
}

func (l *LinkList) AddFront(data Element) {
	head := l.Head
	cur := head.Next
	node := &ListNode{Val: data}
	head.Next = node
	node.Next = cur
}

func (l *LinkList) AddFrontElements(data []Element) {
	for _, v := range data {
		l.AddFront(v)
	}
}

func (l *LinkList) RemoveBack() *ListNode {
	head := l.Head
	cur := &ListNode{}
	for {
		if head.Next.Next != nil {
			head = head.Next
		} else {
			cur = head
			break
		}
	}
	delNode := cur.Next
	cur.Next = nil
	return delNode

}

func (l *LinkList) RemoveFront() *ListNode {
	head := l.Head
	cur := head.Next
	head.Next = cur.Next
	return cur
}

func (l *LinkList) Copy() *LinkList {
	list := NewLinkList()
	head := l.Head
	for head.Next != nil {
		head = head.Next
		list.AddBack(head.Val)
	}
	return list
}

func (n *ListNode) Less(node ListNode) bool {
	if n.Val < node.Val {
		return true
	}
	return false
}

func (n *ListNode) More(node ListNode) bool {
	if n.Val > node.Val {
		return true
	}
	return false
}

func (n *ListNode) Equal(node ListNode) bool {
	if n.Val == node.Val {
		return true
	}
	return false
}

func (l *LinkList) String() string {
	elements := make([]Element, 0)
	head := l.Head
	for head.Next != nil {
		head = head.Next
		elements = append(elements, head.Val)
	}

	res := ""
	for _, v := range elements {
		res += fmt.Sprintf("%v", v) + " "
	}
	return res
}
