package main

import "sort"

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] < intervals[j][0] {
			return true
		}
		return false
	})
	result := make([][]int, 0)
	if len(intervals) < 2 {
		return intervals
	}
	tmp := intervals[0]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] >= tmp[0] && intervals[i][0] <= tmp[1] {
			tmp[1] = max(tmp[1], intervals[i][1])
		} else {
			result = append(result, tmp)
			tmp = intervals[i]
		}
	}
	result = append(result, tmp)
	return result
}
