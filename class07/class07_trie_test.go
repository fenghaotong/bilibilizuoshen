package class07

import (
	"fmt"
	"testing"
)

func TestNewTrie(t *testing.T) {
	trie := NewTrie()
	trie.Insert("abc")
	trie.Insert("ab")
	trie.Insert("cd")
	trie.Insert("abcd")
	trie.Insert("abc")
	fmt.Println(trie.Root.Pass, trie.Root.End)

	fmt.Println(trie.Search("abd"))
	fmt.Println(trie.Search("abc"))

	fmt.Println(trie.PrefixNumber("ab"))

	trie.Delete("abc")
	fmt.Println(trie.Search("abc"))
	fmt.Println(trie.Root.Pass, trie.Root.End)
}

func TestMedianFinder_AddNum(t *testing.T) {
	medianFinder := NewMedianFinder()
	medianFinder.AddNum(1)
	medianFinder.AddNum(2)

	fmt.Println(medianFinder.FindMedian())
	medianFinder.AddNum(3)
	fmt.Println(medianFinder.FindMedian())
	medianFinder.AddNum(4)
	fmt.Println(medianFinder.FindMedian())
	medianFinder.AddNum(5)
	fmt.Println(medianFinder.FindMedian())
	medianFinder.AddNum(6)
	fmt.Println(medianFinder.FindMedian())

	//medianFinder.AddNum(1)
	//medianFinder.AddNum(2)
	//medianFinder.AddNum(3)
	//medianFinder.AddNum(4)
	//medianFinder.AddNum(5)
	////medianFinder.AddNum(6)
	//fmt.Println(medianFinder.FindMedian())
	//medianFinder.PrintPQ()
}

func TestNQueens1(t *testing.T) {
	fmt.Println(NQueens1(14))
}

func TestNQueens2(t *testing.T) {
	fmt.Println(NQueens2(14))
}
