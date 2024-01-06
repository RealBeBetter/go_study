package tree

type TrieNode struct {
	Children map[byte]*TrieNode
	IsEnd    bool
}

type Trie struct {
	Root *TrieNode
}

func Constructor() Trie {
	return Trie{
		Root: &TrieNode{
			Children: make(map[byte]*TrieNode),
			IsEnd:    false,
		},
	}
}

func (trie *Trie) Insert(word string) {
	if len(word) == 0 {
		return
	}

	node := trie.Root
	for i := 0; i < len(word); i++ {
		if node.Children[word[i]] == nil {
			node.Children[word[i]] = &TrieNode{
				Children: make(map[byte]*TrieNode),
				IsEnd:    false,
			}
		}
		node = node.Children[word[i]]
	}

	node.IsEnd = true
}

func (trie *Trie) Search(word string) bool {
	node := trie.Root
	for i := 0; i < len(word); i++ {
		if node.Children[word[i]] == nil {
			return false
		}
		node = node.Children[word[i]]
	}

	return node.IsEnd
}

func (trie *Trie) StartsWith(prefix string) bool {
	node := trie.Root
	for i := 0; i < len(prefix); i++ {
		if node.Children[prefix[i]] == nil {
			return false
		}
		node = node.Children[prefix[i]]
	}

	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
