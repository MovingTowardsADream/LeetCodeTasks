package main

func partitionLabels(s string) []int {
	result := make([]int, 0)
	tmp := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		tmp[s[i]] = i
	}
	left, needIdx := 0, 0
	for right := 0; right < len(s); right++ {
		needIdx = max(needIdx, tmp[s[right]])
		if needIdx <= right {
			result = append(result, (right-left)+1)
			left = right + 1
		}
	}
	return result
}
