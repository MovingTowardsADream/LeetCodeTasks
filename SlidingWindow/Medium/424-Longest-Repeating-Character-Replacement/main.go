package main

func findMax(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	maximum := nums[0]
	for _, num := range nums {
		if num > maximum {
			maximum = num
		}
	}
	return maximum
}

func characterReplacement(s string, k int) int {
	tmp := make([]int, 26)
	left, tmp_max, maximum := 0, 0, 0
	for right := 0; right < len(s); right++ {
		tmp[s[right]-65]++
		tmp_max = max(tmp_max, tmp[s[right]-65])
		if tmp_max+k < right-left+1 {
			tmp[s[left]-65]--
			left++
			tmp_max = findMax(tmp)
		}
		maximum = max(maximum, right-left+1)
	}
	return maximum
}
