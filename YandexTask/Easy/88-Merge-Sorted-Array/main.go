package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	copy(nums1[n:], nums1[:m])
	first, second := n, 0
	for i := 0; i < len(nums1); i++ {
		if second >= n {
			return
		}
		if first < m+n && nums1[first] <= nums2[second] {
			nums1[i] = nums1[first]
			first++
		} else {
			nums1[i] = nums2[second]
			second++
		}
	}
}
