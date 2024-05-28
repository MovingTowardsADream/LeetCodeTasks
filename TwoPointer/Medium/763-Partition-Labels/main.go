package main

import "sort"

func partitionLabels(s string) []int {
	hash := make(map[byte][]int)
	for i := 0; i < len(s); i++ {
		tmp_val := hash[s[i]]
		if len(tmp_val) == 0 {
			hash[s[i]] = []int{i + 1, i + 1}
		} else {
			tmp_val[1] = i + 1
			hash[s[i]] = tmp_val
		}
	}
	intervals := make([][]int, 0)
	for _, interv := range hash {
		intervals = append(intervals, interv)
	}
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] < intervals[j][0] {
			return true
		}
		return false
	})
	result := make([]int, 0)
	tmp := intervals[0]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= tmp[1] && intervals[i][0] >= tmp[0] {
			tmp[1] = max(tmp[1], intervals[i][1])
		} else {
			result = append(result, tmp[1]-tmp[0]+1)
			tmp = intervals[i]
		}
	}
	result = append(result, tmp[1]-tmp[0]+1)
	return result
}
