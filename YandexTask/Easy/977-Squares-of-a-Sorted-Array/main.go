package main

func indexFirstElement(arr []int) (int, int) {
	for i := 1; i < len(arr); i++ {
		if arr[i] >= 0 {
			return i - 1, i
		}
	}
	return len(arr) - 1, len(arr)
}

func sortedSquares(nums []int) []int {
	result := make([]int, 0, len(nums))
	left, right := indexFirstElement(nums)
	for i := 0; i < len(nums); i++ {
		if left < 0 {
			result = append(result, nums[right]*nums[right])
			right++
		} else if right >= len(nums) {
			result = append(result, nums[left]*nums[left])
			left--
		} else {
			if nums[left]*nums[left] <= nums[right]*nums[right] {
				result = append(result, nums[left]*nums[left])
				left--
			} else {
				result = append(result, nums[right]*nums[right])
				right++
			}
		}
	}
	return result
}
