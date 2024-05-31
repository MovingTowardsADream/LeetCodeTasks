package main

func singleNumber(nums []int) int {
	answer := 0
	for i := 0; i < len(nums); i++ {
		answer ^= nums[i]
	}
	return answer
}
