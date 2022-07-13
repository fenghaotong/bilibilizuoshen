package class12

import "container/list"

func GetSlidingWindowMaxArray(arr []int, window int) []int {
	if len(arr) == 0 {
		return []int{}
	}
	dqueue := list.New()
	res := make([]int, len(arr)-window+1)

	if len(arr) <= window {
		for i := 0; i < len(arr); i++ {
			RightMove(arr, i, dqueue)
		}
		v := GetLQueueHead(arr, dqueue)
		res[0] = v
	} else {
		for i := 0; i < window; i++ {
			RightMove(arr, i, dqueue)
		}
		v := GetLQueueHead(arr, dqueue)
		res[0] = v

		for i := window; i < len(arr); i++ {
			RightMove(arr, i, dqueue)
			LeftMove(arr, i-window, dqueue)
			v := GetLQueueHead(arr, dqueue)
			res[i-window+1] = v
		}
	}

	return res
}

// LeftMove
// 将队列左边对应的值删除
func LeftMove(arr []int, i int, dqueue *list.List) {
	back := dqueue.Back()
	value := back.Value.(int)
	if i == value {
		dqueue.Remove(back)
	}
}

// RightMove
// 向队列中加入值，加入之前判断是否具有单调性，如果没有单调性，则将右边的值过期，在进行判断，直到有单调性，将值加入队列
func RightMove(arr []int, i int, dqueue *list.List) {
	for dqueue.Len() > -1 {
		back := dqueue.Back()
		if back != nil {
			value := back.Value.(int)
			if arr[i] <= arr[value] || dqueue.Len() == 0 {
				dqueue.PushBack(i)
				break
			} else {
				dqueue.Remove(back)
				continue
			}
		} else {
			dqueue.PushBack(i)
			break
		}

	}
}

// GetLQueueHead
// 获取队列头的值，并将其过期
func GetLQueueHead(arr []int, dqueue *list.List) int {
	front := dqueue.Front()
	if front != nil {
		dqueue.Remove(front)
		return arr[front.Value.(int)]
	}
	return -1
}
