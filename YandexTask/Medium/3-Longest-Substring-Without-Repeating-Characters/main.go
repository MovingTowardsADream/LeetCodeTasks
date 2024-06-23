package main

func lengthOfLongestSubstring(s string) int {
	tmp := make(map[byte]int, 0)
	left, maximum := 0, 0
	for right := 0; right < len(s); right++ {
		tmp[s[right]]++
		for tmp[s[right]] > 1 {
			tmp[s[left]]--
			left++
		}
		maximum = max(maximum, right-left+1)
	}
	return maximum
}
