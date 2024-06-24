package main

func permutation(result *[][]int, nums []int, n int) {
	tmp := make([]int, len(nums))
	if n < 2 {
		copy(tmp, nums)
		*result = append(*result, tmp)
	} else {
		for i := n - 1; i >= 0; i-- {
			nums[i], nums[n-1] = nums[n-1], nums[i]
			permutation(result, nums, n-1)
			nums[i], nums[n-1] = nums[n-1], nums[i]
		}
	}
}

func permute(nums []int) [][]int {
	var result [][]int
	permutation(&result, nums, len(nums))
	return result
}
