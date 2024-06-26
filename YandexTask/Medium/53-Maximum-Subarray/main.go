package main

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	local_max, global_max := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if local_max < 0 {
			local_max = 0
		}
		local_max += nums[i]
		global_max = max(global_max, local_max)
	}
	return global_max
}
