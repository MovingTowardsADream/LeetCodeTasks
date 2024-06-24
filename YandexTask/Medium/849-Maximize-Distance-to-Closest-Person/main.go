package main

func maxDistToClosest(seats []int) int {
	left, maximum := 0, 0
	for right := 0; right < len(seats); right++ {
		if seats[right] != 0 {
			left = right
		}
		maximum = max(maximum, right-left)
	}
	maximum = (maximum + 1) / 2

	maxOnLeft, maxOnRight := 0, 0
	for seats[maxOnLeft] != 1 {
		maxOnLeft++
	}
	for seats[len(seats)-maxOnRight-1] != 1 {
		maxOnRight++
	}
	return max(maximum, max(maxOnRight, maxOnLeft))
}
