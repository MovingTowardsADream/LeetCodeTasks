package main

import "fmt"

// O(n+m)
// Алгоритм интересен, если используется много похожих, но не совпадающих слов

// Заполнение количества суффиксов и преффиксов
func prefixFunction(s string) []int {
	pref := make([]int, len(s))
	for i := 1; i < len(s); i++ {
		cur := pref[i-1]
		for s[i] != s[cur] && cur > 0 {
			cur = pref[cur-1]
		}
		if s[i] == s[cur] {
			pref[i] = cur + 1
		}
	}
	return pref
}

func knuthMorrisPratt(all, word string) int {
	tmp := prefixFunction(word)
	first, second := 0, 0
	for first < len(all) && second < len(word) {
		if all[first] == word[second] {
			first++
			second++
		} else {
			if second == 0 {
				first++
			} else {
				second = tmp[second-1]
			}
		}
	}
	if len(word) == second {
		return first - second
	}
	return -1
}

func main() {
	fmt.Println(knuthMorrisPratt("abcabeabcabcabd", "abcabd"))
}
