package class08

import "algorithm/datastruct"

// 不使用额外控制逆序栈

func ReverseStack(stack *datastruct.Stack) {
	if stack.Empty() {
		return
	}

	i := getStackBottom(stack)
	ReverseStack(stack)
	stack.Push(i)
}

// 获取栈底元素
func getStackBottom(stack *datastruct.Stack) interface{} {
	pop := stack.Pop()
	if stack.Empty() {
		return pop
	} else {
		last := getStackBottom(stack)
		stack.Push(pop)
		return last
	}
}
