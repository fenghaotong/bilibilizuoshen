package class10

import "algorithm/datastruct"

type Element struct {
	Value int
}

type UnionFindSet struct {
	ElementMap map[int]Element
	FatherMap  map[Element]Element
	SizeMap    map[Element]int
}

func NewUnionFindSet(arr []int) *UnionFindSet {
	unionSet := &UnionFindSet{
		ElementMap: make(map[int]Element, 0),
		FatherMap:  make(map[Element]Element, 0),
		SizeMap:    make(map[Element]int, 0),
	}
	for _, v := range arr {
		element := Element{v}
		unionSet.ElementMap[v] = element
		unionSet.FatherMap[element] = element
		unionSet.SizeMap[element] = 1
	}
	return unionSet
}

func (u *UnionFindSet) ContainKey(key int) bool {
	for k, _ := range u.ElementMap {
		if k == key {
			return true
		}
	}
	return false
}

func (u *UnionFindSet) FindHead(e Element) Element {
	stack := datastruct.NewStack()
	for e != u.FatherMap[e] {
		stack.Push(e)
		e = u.FatherMap[e]
	}

	for !stack.Empty() {
		pop := stack.Pop().(Element)
		u.FatherMap[pop] = e
	}
	return e
}

func (u *UnionFindSet) IsSameSet(a, b int) bool {
	if u.ContainKey(a) && u.ContainKey(b) {
		return u.FindHead(u.ElementMap[a]) == u.FindHead(u.ElementMap[b])
	}
	return false
}

func (u *UnionFindSet) Union(a, b int) {
	if u.ContainKey(a) && u.ContainKey(b) {
		aF := u.FindHead(u.ElementMap[a])
		bF := u.FindHead(u.ElementMap[b])

		if aF != bF {
			big := Element{}
			small := Element{}
			if u.SizeMap[aF] >= u.SizeMap[bF] {
				big = aF
				small = bF
			} else {
				big = bF
				small = aF
			}
			u.FatherMap[small] = big
			u.SizeMap[big] = u.SizeMap[aF] + u.SizeMap[bF]
			delete(u.SizeMap, small)
		}
	}
}
