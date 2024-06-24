package main

func findClosestIndex(arr []int, value int) int {
	index := 0
	for i := 1; i < len(arr); i++ {
		index++
		if arr[i] == value {
			break
		}
		if arr[i] > value {
			if arr[i]-value >= value-arr[i-1] {
				index--
			}
			break
		}
	}
	return index
}

func findClosestElements(arr []int, k int, x int) []int {
	left := findClosestIndex(arr, x)
	right := left
	for i := 1; i < k; i++ {
		if left == 0 {
			right++
		} else if right == len(arr)-1 {
			left--
		} else if arr[right+1]-x >= x-arr[left-1] {
			left--
		} else {
			right++
		}
	}
	return arr[left : right+1]
}
