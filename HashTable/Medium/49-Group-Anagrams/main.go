package main

import "sort"

func groupAnagrams(strs []string) [][]string {
	dict := make(map[string][]string)
	for i := 0; i < len(strs); i++ {
		tmp := []rune(strs[i])
		sort.Slice(tmp, func(i, j int) bool {
			return tmp[i] < tmp[j]
		})
		dict[string(tmp)] = append(dict[string(tmp)], strs[i])
	}
	result := make([][]string, len(dict))
	index := 0
	for _, value := range dict {
		result[index] = value
		index++
	}
	return result
}
