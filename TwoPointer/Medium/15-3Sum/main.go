package main

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result := make([][]int, 0)
	var left, right int
	for i := 0; i < len(nums)-2; i++ {
		if nums[i]+nums[i+1]+nums[i+2] > 0 {
			break
		}
		if nums[i]+nums[len(nums)-1]+nums[len(nums)-2] < 0 {
			continue
		}
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right = i+1, len(nums)-1
		for left < right {
			if nums[i]+nums[left]+nums[right] == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				left++
				for nums[left] == nums[left-1] && left < right {
					left++
				}
				right--
				for nums[right] == nums[right+1] && left < right {
					right--
				}
			} else if nums[i]+nums[left]+nums[right] > 0 {
				right--
			} else {
				left++
			}
		}
	}
	return result
}
