package main

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	var global_max, local_max int
	for left < right {
		local_max = (right - left) * min(height[left], height[right])
		if local_max > global_max {
			global_max = local_max
		}
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return global_max
}
