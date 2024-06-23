package main

func longestOnes(nums []int, k int) int {
	left, amount_zero, maximum := 0, 0, 0
	for right := 0; right < len(nums); right++ {
		if nums[right] == 0 {
			if amount_zero >= k {
				for nums[left] != 0 {
					left++
				}
				left++
			}
			amount_zero++
		}
		maximum = max(maximum, right-left+1)
	}
	return maximum
}
