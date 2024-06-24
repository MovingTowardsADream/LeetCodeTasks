package main

func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	first := 0
	for second := 0; second < len(t); second++ {
		if s[first] == t[second] {
			first++
		}
		if first >= len(s) {
			return true
		}
	}
	return false
}
