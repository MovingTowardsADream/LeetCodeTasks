package main

import "fmt"

func shortest_distance(graph [][]int, begin int) map[int]int {
	result := map[int]int{0: 0}
	set := map[int]struct{}{begin: {}}
	queue := []int{begin}
	for i := 0; len(queue) > 0; i++ {
		point := queue[0]
		queue = queue[1:]
		for _, val := range graph[point] {
			if _, ok := result[val]; !ok {
				result[val] = result[point] + 1
			} else {
				result[val] = min(result[val], result[point]+1)
			}
			if _, ok := set[val]; !ok {
				set[val] = struct{}{}
				queue = append(queue, val)
			}
		}
	}
	return result
}

func main() {
	graph := [][]int{{1, 4, 5}, {0, 2, 6}, {1, 3, 7}, {2, 4, 8}, {0, 3, 9},
		{0, 7, 8}, {1, 8, 9}, {2, 5, 9}, {3, 5, 6}, {4, 6, 7}}
	fmt.Println(shortest_distance(graph, 0))
}
