package class04

import (
	"algorithm/datastruct"
	"fmt"
	"testing"
)

func TestHashTable(t *testing.T) {
	HashTable()
}

func TestIsPalindromic1(t *testing.T) {
	list := datastruct.NewLinkList()
	elements := []datastruct.Element{1, 2, 2, 2, 1}
	list.AddBackElements(elements)

	fmt.Println(IsPalindromic1(list))
}

func TestIsPalindromic2(t *testing.T) {
	list := datastruct.NewLinkList()
	elements := []datastruct.Element{1, 2, 2, 2, 1}
	list.AddBackElements(elements)

	fmt.Println(IsPalindromic2(list))
}

func TestIsPalindromic3(t *testing.T) {
	list := datastruct.NewLinkList()
	elements := []datastruct.Element{1, 3, 3, 2, 1}
	list.AddBackElements(elements)

	fmt.Println(IsPalindromic3(list))
}

func TestListPartition1(t *testing.T) {
	list := datastruct.NewLinkList()
	elements := []datastruct.Element{1, 3, 3, 2, 1, 4, 6, 7, 2}
	list.AddBackElements(elements)
	fmt.Println(ListPartition1(list, 7))
	fmt.Println(ListPartition1(list, 3))
	fmt.Println(ListPartition1(list, 5))
}

func TestListPartition2(t *testing.T) {
	list := datastruct.NewLinkList()
	elements := []datastruct.Element{1, 3, 3, 2, 1, 4, 6, 7, 2}
	list.AddBackElements(elements)

	fmt.Println(ListPartition2(list, 7))
	fmt.Println(ListPartition2(list, 3))
	fmt.Println(ListPartition2(list, 5))
}

func TestGetLoopNode(t *testing.T) {
	// 1->2->3->4->5->6->7->null
	head1 := &datastruct.ListNode{Val: 1}
	head1.Next = &datastruct.ListNode{Val: 2}
	head1.Next.Next = &datastruct.ListNode{Val: 3}
	head1.Next.Next.Next = &datastruct.ListNode{Val: 4}
	head1.Next.Next.Next.Next = &datastruct.ListNode{Val: 5}
	head1.Next.Next.Next.Next.Next = &datastruct.ListNode{Val: 6}
	head1.Next.Next.Next.Next.Next.Next = &datastruct.ListNode{Val: 7}

	// 0->9->8->6->7->null
	head2 := &datastruct.ListNode{Val: 0}
	head2.Next = &datastruct.ListNode{Val: 9}
	head2.Next.Next = &datastruct.ListNode{Val: 8}
	head2.Next.Next.Next = head1.Next.Next.Next.Next.Next // 8->6

	list1 := datastruct.NewLinkList()
	list2 := datastruct.NewLinkList()
	list1.Head = head1
	list2.Head = head2
	fmt.Println(GetIntersectNode(list1, list2).Val)

	head1 = nil
	head2 = nil

	// 1->2->3->4->5->6->7->4...
	head1 = &datastruct.ListNode{Val: 1}
	head1.Next = &datastruct.ListNode{Val: 2}
	head1.Next.Next = &datastruct.ListNode{Val: 3}
	head1.Next.Next.Next = &datastruct.ListNode{Val: 4}
	head1.Next.Next.Next.Next = &datastruct.ListNode{Val: 5}
	head1.Next.Next.Next.Next.Next = &datastruct.ListNode{Val: 6}
	head1.Next.Next.Next.Next.Next.Next = &datastruct.ListNode{Val: 7}
	head1.Next.Next.Next.Next.Next.Next = head1.Next.Next.Next // 7->4

	// 0->9->8->2...
	head2 = &datastruct.ListNode{Val: 0}
	head2.Next = &datastruct.ListNode{Val: 9}
	head2.Next.Next = &datastruct.ListNode{Val: 8}
	head2.Next.Next.Next = head1.Next // 8->2
	list1 = datastruct.NewLinkList()
	list2 = datastruct.NewLinkList()
	list1.Head = head1
	list2.Head = head2
	fmt.Println(GetIntersectNode(list1, list2).Val)

	head2 = nil

	// 0->9->8->6->7->4->5->6..
	head2 = &datastruct.ListNode{Val: 0}
	head2.Next = &datastruct.ListNode{Val: 9}
	head2.Next.Next = &datastruct.ListNode{Val: 8}
	head2.Next.Next.Next = head1.Next.Next.Next.Next.Next // 8->6
	list1 = datastruct.NewLinkList()
	list2 = datastruct.NewLinkList()
	list1.Head = head1
	list2.Head = head2
	fmt.Println(GetIntersectNode(list1, list2).Val)
}
