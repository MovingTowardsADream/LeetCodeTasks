package main

func lengthOfLongestSubstring(s string) int {
	left, maximum := 0, 0
	tmp := make(map[byte]int)
	for right := 0; right < len(s); right++ {
		if tmp[s[right]] > 0 {
			left = max(left, tmp[s[right]])
		}
		tmp[s[right]] = right + 1
		maximum = max(maximum, right-left+1)
	}
	return maximum
}
