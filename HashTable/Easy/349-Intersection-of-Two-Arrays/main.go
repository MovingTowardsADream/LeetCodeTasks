package main

func intersection(nums1 []int, nums2 []int) []int {
	tmp := make(map[int]int)
	for i := 0; i < len(nums1); i++ {
		tmp[nums1[i]] = 1
	}
	for i := 0; i < len(nums2); i++ {
		if tmp[nums2[i]] > 0 {
			tmp[nums2[i]] = 2
		}
	}
	result := make([]int, 0)
	for key, value := range tmp {
		if value > 1 {
			result = append(result, key)
		}
	}
	return result
}
