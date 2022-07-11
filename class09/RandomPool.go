package class09

import "math/rand"

// 为了保证获取随机数字的时间复杂度为O(1)

type RandomPool struct {
	Keys        []string
	KeyIndexMap map[string]int
	IndexKeyMap map[int]string
	Size        int
}

func NewRandomPool() *RandomPool {
	return &RandomPool{
		Keys:        make([]string, 0),
		KeyIndexMap: make(map[string]int, 0),
		IndexKeyMap: make(map[int]string, 0),
		Size:        0,
	}
}

func (RP *RandomPool) Contains(key string) bool {
	for k, _ := range RP.KeyIndexMap {
		if k == key {
			return true
		}
	}
	return false
}

func (RP *RandomPool) Insert(str string) {
	if !RP.Contains(str) {
		RP.KeyIndexMap[str] = RP.Size
		RP.IndexKeyMap[RP.Size] = str
		RP.Size++
	}
}

func (RP *RandomPool) Delete(str string) {
	if RP.Contains(str) {
		RP.Size--
		deleteIndex := RP.KeyIndexMap[str]
		lastIndex := RP.Size
		lastKey := RP.IndexKeyMap[lastIndex]
		RP.KeyIndexMap[lastKey] = deleteIndex
		RP.IndexKeyMap[deleteIndex] = lastKey
		delete(RP.KeyIndexMap, str)
		delete(RP.IndexKeyMap, lastIndex)
	}

}

func (RP *RandomPool) GetRandom() string {
	if RP.Size == 0 {
		return ""
	}

	rand.Seed(rand.Int63())
	randomIndex := rand.Intn(RP.Size)
	return RP.IndexKeyMap[randomIndex]
}
