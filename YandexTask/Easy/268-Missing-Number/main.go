package main

func missingNumber(nums []int) int {
	res := 0
	for i := 0; i < len(nums); i++ {
		res ^= i
		res ^= nums[i]
	}
	return res ^ len(nums)
}
