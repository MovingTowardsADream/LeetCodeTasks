package main

// len(nums) > 1
func searchTwoMax(nums []int) (int, int) {
	first, second := 0, 1
	if nums[first] < nums[second] {
		first, second = second, first
	}
	for i := 2; i < len(nums); i++ {
		if nums[i] > nums[second] {
			if nums[i] > nums[first] {
				first, second = i, first
			} else {
				second = i
			}
		}
	}
	return first, second
}

func lastStoneWeight(stones []int) int {
	var first, second, tmp int
	for len(stones) > 1 {
		first, second = searchTwoMax(stones)
		tmp = max(stones[first]-stones[second], stones[second]-stones[first])
		if first > second {
			first, second = second, first
		}
		stones = append(stones[:second], stones[second+1:]...)
		if tmp == 0 {
			stones = append(stones[:first], stones[first+1:]...)
		} else {
			stones[first] = tmp
		}
	}
	if len(stones) == 0 {
		return 0
	}
	return stones[0]
}
