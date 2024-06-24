package main

func intersect(nums1 []int, nums2 []int) []int {
	hash := make(map[int]int)
	for i := 0; i < len(nums1); i++ {
		hash[nums1[i]]++
	}
	result := make([]int, 0)
	for i := 0; i < len(nums2); i++ {
		if hash[nums2[i]] > 0 {
			hash[nums2[i]]--
			result = append(result, nums2[i])
		}
	}
	return result
}
