package class09

import (
	"fmt"
	"testing"
)

func TestNewRandomPool(t *testing.T) {
	pool := NewRandomPool()
	pool.Insert("abc")
	pool.Insert("a")
	pool.Insert("bc")
	for i := 0; i < 10; i++ {
		fmt.Println(pool.GetRandom())
	}

}
