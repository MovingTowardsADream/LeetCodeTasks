package main

func twoSum(nums []int, target int) []int {
	hash := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if hash[target-nums[i]] > 0 {
			return []int{hash[target-nums[i]] - 1, i}
		}
		hash[nums[i]] = i + 1
	}
	return []int{-1, -1}
}
