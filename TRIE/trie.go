package main

type Node struct {
	links [26]*Node
	flag  bool
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
		flag: false,
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
	}
	currentNode.flag = true
}

func (t *TRIE) Search(word string) bool {
	currentNode := t.root
	n := len(word)
	for i := 0; i < n; i++ {
		if !currentNode.ContainsKey(rune(word[i])) {
			return false
		}
		currentNode = currentNode.Get(rune(word[i]))
	}
	return currentNode.flag
}

func (t *TRIE) StartsWith(word string) bool {

	currentNode := t.root
	n := len(word)

	for i := 0; i < n; i++ {
		if !currentNode.ContainsKey(rune(word[i])) {
			return false
		}
		currentNode = currentNode.Get(rune(word[i]))
	}
	return true
}
