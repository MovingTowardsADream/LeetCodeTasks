package main

import (
	"fmt"
	"sort"
	"strconv"
)

func summaryRanges(nums []int) []string {
	sort.Ints(nums)
	result := make([]string, 0)
	if len(nums) == 0 {
		return result
	}
	left, right := 0, 0
	for i := 1; i < len(nums); i++ {
		if nums[right] == nums[i]-1 {
			right++
		} else {
			if left == right {
				result = append(result, strconv.Itoa(nums[left]))
			} else {
				result = append(result, fmt.Sprintf("%s->%s", strconv.Itoa(nums[left]), strconv.Itoa(nums[right])))
			}
			left, right = i, i
		}
	}
	if left == right {
		result = append(result, strconv.Itoa(nums[left]))
	} else {
		result = append(result, fmt.Sprintf("%s->%s", strconv.Itoa(nums[left]), strconv.Itoa(nums[right])))
	}
	return result
}
