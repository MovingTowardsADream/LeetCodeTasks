package main

func peakIndexInMountainArray(arr []int) int {
	left, right := 0, len(arr)-1
	for left <= right {
		mid := (left + right) / 2
		if arr[mid] > arr[mid+1] && arr[mid] > arr[mid-1] {
			return mid
		} else if mid == 0 || arr[mid] > arr[mid-1] {
			left = mid + 1
		} else if mid == len(arr) || arr[mid] > arr[mid+1] {
			right = mid - 1
		}
	}
	return -1
}
