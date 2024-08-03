package main

import "fmt"

func bfs(graph [][]int, begin int) []int {
	result := make([]int, 0)
	set := map[int]struct{}{begin: {}}
	queue := []int{begin}
	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]
		result = append(result, u)
		for _, value := range graph[u] {
			if _, ok := set[value]; !ok {
				queue = append(queue, value)
				set[value] = struct{}{}
			}
		}
	}
	return result
}

func main() {
	graph := [][]int{{1, 4, 5}, {0, 2, 6}, {1, 3, 7}, {2, 4, 8}, {0, 3, 9},
		{0, 7, 8}, {1, 8, 9}, {2, 5, 9}, {3, 5, 6}, {4, 6, 7}}
	fmt.Println(bfs(graph, 0))
}
