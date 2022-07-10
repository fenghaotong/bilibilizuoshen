package datastruct

import (
	"fmt"
	"testing"
)

func TestContainsKey(t *testing.T) {
	m := make(map[int]GraphNode, 0)
	m[0] = GraphNode{Val: 0}
	m[1] = GraphNode{Val: 1}
	m[2] = GraphNode{Val: 2}
	fmt.Println(ContainsKey(m, 0))
	fmt.Println(ContainsKey(m, 3))
}
