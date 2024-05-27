package main

func searchMatrix(matrix [][]int, target int) bool {
	n, m := len(matrix), len(matrix[0])
	left, right := 0, n-1
	if matrix[0][0] > target || matrix[n-1][m-1] < target {
		return false
	}
	var line int
	for left <= right {
		mid := (left + right) / 2
		if matrix[mid][0] <= target && matrix[mid][m-1] >= target {
			line = mid
			break
		} else if matrix[mid][0] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	left, right = 0, m-1
	for left <= right {
		mid := (left + right) / 2
		if matrix[line][mid] == target {
			return true
		} else if matrix[line][mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return false
}
