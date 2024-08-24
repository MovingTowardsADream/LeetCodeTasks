package main

import "fmt"

func dfs(graph [][]int, begin int) []int {
	result := []int{}
	set := map[int]struct{}{begin: {}}
	stack := []int{begin}
	for len(stack) > 0 {
		v := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, v)
		for _, val := range graph[v] {
			if _, ok := set[val]; !ok {
				set[val] = struct{}{}
				stack = append(stack, val)
			}
		}
	}
	return result
}

func main() {
	graph := [][]int{{1, 4, 5}, {0, 2, 6}, {1, 3, 7}, {2, 4, 8}, {0, 3, 9},
		{0, 7, 8}, {1, 8, 9}, {2, 5, 9}, {3, 5, 6}, {4, 6, 7}}
	fmt.Println(dfs(graph, 0))
}
