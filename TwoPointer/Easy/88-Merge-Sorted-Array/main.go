package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}
	second, three := n, 0
	copy(nums1[n:], nums1[:m])
	for i := 0; i < n; i++ {
		nums1[i] = 0
	}
	for first := 0; first < n+m; first++ {
		if (three >= n) || (second < n+m && nums1[second] <= nums2[three]) {
			nums1[first] = nums1[second]
			second++
		} else {
			nums1[first] = nums2[three]
			three++
		}
	}
}
