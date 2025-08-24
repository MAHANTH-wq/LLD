package main

type Node struct {
	links           [26]*Node
	endsWithCount   int
	prefixWithCount int
}

func (n *Node) ContainsKey(w rune) bool {

	return n.links[w-'a'] != nil

}

func (n *Node) Get(w rune) *Node {
	if n.links[w-'a'] != nil {
		return n.links[w-'a']
	}
	return nil
}

func (n *Node) Put(w rune) {
	n.links[w-'a'] = newNode()
}

func newNode() *Node {
	return &Node{
		endsWithCount:   0,
		prefixWithCount: 0,
	}
}

type TRIE struct {
	root *Node
}

func NewTRIE() *TRIE {
	return &TRIE{
		root: newNode(),
	}
}

func (t *TRIE) Insert(word string) {

	currentNode := t.root
	n := len(word)

	for i := 0; i < n; i++ {

		if !currentNode.ContainsKey(rune(word[i])) {
			currentNode.Put(rune(word[i]))
		}
		currentNode = currentNode.Get(rune(word[i]))
		currentNode.prefixWithCount++
	}
	currentNode.endsWithCount++
}

func (t *TRIE) Search(word string) int {
	currentNode := t.root
	n := len(word)
	for i := 0; i < n; i++ {
		if !currentNode.ContainsKey(rune(word[i])) {
			return 0
		}
		currentNode = currentNode.Get(rune(word[i]))
	}
	return currentNode.endsWithCount
}

func (t *TRIE) StartsWith(word string) int {

	currentNode := t.root
	n := len(word)

	for i := 0; i < n; i++ {
		if !currentNode.ContainsKey(rune(word[i])) {
			return 0
		}
		currentNode = currentNode.Get(rune(word[i]))
	}
	return currentNode.prefixWithCount
}

func (t *TRIE) LongestCompleteString(root *Node) string {

	maxString := ""
	for index, value := range root.links {
		if value != nil && value.endsWithCount > 0 {
			ll := t.LongestCompleteString(value)
			if len(ll) > len(maxString) || maxString == "" {
				maxString = string(rune('a'+index)) + ll
			}
		}
	}
	return maxString
}

func CountNoOfDistinctSubStrings(word string) int {

	trie := NewTRIE()
	n := len(word)
	count := 0
	for i := 0; i < n; i++ {
		currentNode := trie.root
		for j := i; j < n; j++ {
			if !currentNode.ContainsKey(rune(word[j])) {
				count++
				currentNode.Put(rune(word[j]))
			}
			currentNode = currentNode.Get(rune(word[j]))
		}
	}

	return (count + 1)

}
