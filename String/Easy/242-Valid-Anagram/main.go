package main

func isAnagram(s string, t string) bool {
	dict := make(map[rune]int)
	for _, i := range s {
		dict[i]++
	}
	for _, i := range t {
		if dict[i] <= 0 {
			return false
		}
		dict[i]--
	}
	for _, value := range dict {
		if value != 0 {
			return false
		}
	}
	return true
}
