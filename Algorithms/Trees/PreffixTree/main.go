package main

import "fmt"

type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
}

type Trie struct {
	root *TrieNode
}

func NewTrie() *Trie {
	return &Trie{root: &TrieNode{children: make(map[rune]*TrieNode)}}
}

func (t *Trie) Insert(word string) {
	node := t.root
	for _, r := range word {
		if _, ok := node.children[r]; !ok {
			node.children[r] = &TrieNode{children: make(map[rune]*TrieNode)}
		}
		node = node.children[r]
	}
	node.isEnd = true
}

func (t *Trie) Search(word string) bool {
	node := t.root
	for _, r := range word {
		if _, ok := node.children[r]; !ok {
			return false
		}
		node = node.children[r]
	}
	return node.isEnd
}

func (t *Trie) StartsWith(prefix string) bool {
	node := t.root
	for _, r := range prefix {
		if _, ok := node.children[r]; !ok {
			return false
		}
		node = node.children[r]
	}
	return true
}

func main() {
	trie := NewTrie()
	trie.Insert("apple")
	fmt.Println(trie.Search("apple")) // returns true
	fmt.Println(trie.Search("app"))   // returns false
}
