package main

func isMonotonic(nums []int) bool {
	if len(nums) < 2 {
		return true
	}
	if nums[0] <= nums[len(nums)-1] {
		for i := 1; i < len(nums); i++ {
			if nums[i] < nums[i-1] {
				return false
			}
		}
	} else {
		for i := 1; i < len(nums); i++ {
			if nums[i] > nums[i-1] {
				return false
			}
		}
	}
	return true
}
