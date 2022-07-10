package class07

// 前缀树

type TrieNode struct {
	Pass  int
	End   int
	Nexts []*TrieNode
}

func NewTrieNode() *TrieNode {
	return &TrieNode{
		Pass:  0,
		End:   0,
		Nexts: make([]*TrieNode, 26),
	}
}

type Trie struct {
	Root *TrieNode
}

func NewTrie() *Trie {
	return &Trie{
		Root: NewTrieNode(),
	}
}

func (t *Trie) Insert(word string) {
	if len(word) == 0 {
		return
	}
	root := t.Root
	root.Pass++
	for _, v := range word {
		node := NewTrieNode()
		index := int(v) - 65 - 32
		if root.Nexts[index] == nil {
			root.Nexts[index] = node
		}
		root = root.Nexts[index]
		root.Pass++
	}
	root.End++
}

func (t *Trie) Search(word string) int {
	if len(word) == 0 {
		return 0
	}

	root := t.Root
	for _, v := range word {
		index := int(v) - 65 - 32
		if root.Nexts[index] == nil {
			return 0
		}
		root = root.Nexts[index]
	}
	return root.End
}

func (t *Trie) Delete(word string) {
	if t.Search(word) != 0 {
		root := t.Root
		root.Pass--
		for _, v := range word {
			index := int(v) - 65 - 32
			node := root.Nexts[index]
			node.Pass = node.Pass - 1
			if node.Pass == 0 {
				node = nil
				return
			}
			root = root.Nexts[index]
		}
		root.End--
	}

}

func (t *Trie) PrefixNumber(pre string) int {
	if len(pre) == 0 {
		return 0
	}
	root := t.Root
	for _, v := range pre {
		index := int(v) - 65 - 32
		if root.Nexts[index] == nil {
			return 0
		}
		root = root.Nexts[index]
	}
	return root.Pass
}
