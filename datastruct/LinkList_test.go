package datastruct

import (
	"fmt"
	"testing"
)

func TestListNode_Init(t *testing.T) {
	node := NewLinkList()
	node.AddBack(Element(1))
	node.AddBack(Element(2))
	node.AddBack(Element(3))
	node.AddBack(Element(5))
	node.AddFront(Element(7))

	fmt.Println(node)

	fmt.Println(node.RemoveBack())

	fmt.Println(node)

	fmt.Println(node.RemoveFront())
	fmt.Println(node)
}
