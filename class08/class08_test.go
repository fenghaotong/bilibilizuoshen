package class08

import (
	"algorithm/datastruct"
	"fmt"
	"testing"
)

func TestHanoi(t *testing.T) {
	Hanoi(3)
}

func TestPrintStringSubSeq(t *testing.T) {
	PrintStringSubSeq("abc")
}

func TestPrintStringPermutations(t *testing.T) {
	PrintStringPermutations("abc")
}

func TestWin(t *testing.T) {
	arr := []int{1, 2, 100, 4}
	fmt.Println(Win(arr))
}

func TestReverseStack(t *testing.T) {
	stack := datastruct.NewStack()
	stack.Push(3)
	stack.Push(2)
	stack.Push(1)
	//fmt.Println(stack.Pop())
	//fmt.Println(stack.Pop())
	//fmt.Println(stack.Pop())
	ReverseStack(stack)
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
}

func TestConvert(t *testing.T) {
	arr := []uint8{'1', '2', '2', '1'}
	fmt.Println(Convert(arr))
}

func TestGetMaxValue(t *testing.T) {
	weights := []int{10, 5, 4, 2, 3}
	values := []int{20, 10, 15, 3, 1}
	fmt.Println(GetMaxValue(weights, values, 11))
}
