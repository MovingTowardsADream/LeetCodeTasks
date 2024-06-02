package main

import "sort"

func topKFrequent(nums []int, k int) []int {
	tmp := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		tmp[nums[i]]++
	}
	result := make([][]int, len(tmp))
	index := 0
	for key, value := range tmp {
		result[index] = []int{key, value}
		index++
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i][1] > result[j][1]
	})
	res := make([]int, k)
	for i := 0; i < min(k, len(result)); i++ {
		res[i] = result[i][0]
	}
	return res
}
