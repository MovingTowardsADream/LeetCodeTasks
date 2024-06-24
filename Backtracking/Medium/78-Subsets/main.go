package main

func permutation(result *[][]int, nums, arr []int, lenght int, index int) {
	if len(arr) == lenght {
		tmp := make([]int, lenght)
		copy(tmp, arr)
		*result = append(*result, tmp)
	} else if len(nums)-index < lenght-len(arr) {
		return
	} else {
		for i := index; i < len(nums); i++ {
			arr = append(arr, nums[i])
			permutation(result, nums, arr, lenght, i+1)
			arr = arr[:len(arr)-1]
		}
	}
}

func subsets(nums []int) [][]int {
	var result [][]int
	result = append(result, []int{})
	for i := 1; i < len(nums)+1; i++ {
		tmp := make([]int, 0, i)
		permutation(&result, nums, tmp, i, 0)
	}
	return result
}
