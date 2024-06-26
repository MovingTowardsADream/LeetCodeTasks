package main

func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	result := make([][]int, 0)
	first, second := 0, 0
	for first < len(firstList) && second < len(secondList) {
		if firstList[first][1] >= secondList[second][1] {
			if firstList[first][0] <= secondList[second][1] {
				if secondList[second][0] >= firstList[first][0] {
					result = append(result, []int{secondList[second][0], secondList[second][1]})
				} else {
					result = append(result, []int{firstList[first][0], secondList[second][1]})
				}
			}
			second++
		} else {
			if secondList[second][0] <= firstList[first][1] {
				if firstList[first][0] >= secondList[second][0] {
					result = append(result, []int{firstList[first][0], firstList[first][1]})
				} else {
					result = append(result, []int{secondList[second][0], firstList[first][1]})
				}
			}
			first++
		}
	}
	return result
}
