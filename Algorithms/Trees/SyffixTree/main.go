package main

import "fmt"

// SuffixTreeNode представляет узел суффиксного дерева
type SuffixTreeNode struct {
	children map[rune]*SuffixTreeNode
	isEnd    bool
}

// SuffixTree представляет само суффиксное дерево
type SuffixTree struct {
	root *SuffixTreeNode
}

// NewSuffixTree инициализирует новое суффиксное дерево
func NewSuffixTree() *SuffixTree {
	return &SuffixTree{
		root: &SuffixTreeNode{children: make(map[rune]*SuffixTreeNode)},
	}
}

// Insert добавляет суффикс в суффиксное дерево
func (tree *SuffixTree) Insert(suffix string) {
	node := tree.root
	for _, char := range suffix {
		if _, exists := node.children[char]; !exists {
			node.children[char] = &SuffixTreeNode{children: make(map[rune]*SuffixTreeNode)}
		}
		node = node.children[char]
	}
	node.isEnd = true
}

// Build строит суффиксное дерево для заданной строки
func (tree *SuffixTree) Build(s string) {
	for i := range s {
		tree.Insert(s[i:])
	}
}

// Search ищет подстроку в суффиксном дереве и возвращает true, если подстрока найдена
func (tree *SuffixTree) Search(substring string) bool {
	node := tree.root
	for _, char := range substring {
		if _, exists := node.children[char]; !exists {
			return false
		}
		node = node.children[char]
	}
	return true
}

func main() {
	tree := NewSuffixTree()
	s := "banana"
	tree.Build(s)

	fmt.Println(tree.Search("anan"))   // true
	fmt.Println(tree.Search("nana"))   // true
	fmt.Println(tree.Search("banana")) // true
	fmt.Println(tree.Search("apple"))  // false
}
