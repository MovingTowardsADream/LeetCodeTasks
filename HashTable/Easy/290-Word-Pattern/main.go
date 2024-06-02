package main

import "strings"

func wordPattern(pattern string, s string) bool {
	words := strings.Split(s, " ")
	if len(words) != len(pattern) {
		return false
	}
	letter, word := make(map[byte]int), make(map[string]int)
	for i := 0; i < len(pattern); i++ {
		if letter[pattern[i]] != word[words[i]] {
			return false
		}
		letter[pattern[i]], word[words[i]] = i+1, i+1
	}
	return true
}
