package main

func longestSubarray(nums []int) int {
	left, amount, global_max := 0, 0, 0
	for right := 0; right < len(nums); right++ {
		if nums[right] == 0 {
			if amount == 0 {
				amount++
			} else {
				for nums[left] != 0 {
					left++
				}
				left++
			}
		}
		global_max = max(global_max, right-left)
	}
	return global_max
}
