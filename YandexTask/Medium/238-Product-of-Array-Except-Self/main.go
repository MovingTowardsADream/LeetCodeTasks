package main

func productExceptSelf(nums []int) []int {
	result := make([]int, len(nums))
	if len(nums) == 0 {
		return result
	}
	tmp := nums[0]
	for i := 1; i < len(nums); i++ {
		result[i] = tmp
		tmp *= nums[i]
	}
	tmp = nums[len(nums)-1]
	for i := len(nums) - 2; i > 0; i-- {
		result[i] *= tmp
		tmp *= nums[i]
	}
	if len(nums) > 1 {
		result[0] = tmp
	}
	return result
}
