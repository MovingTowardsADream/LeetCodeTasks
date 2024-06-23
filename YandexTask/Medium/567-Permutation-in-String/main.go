package main

func sumZero(nums [26]int) bool {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	return sum == 0
}

func checkInclusion(s1 string, s2 string) bool {
	var tmp [26]int
	for i := 0; i < len(s1); i++ {
		tmp[s1[i]-97]++
	}
	left := 0
	for right := 0; right < len(s2); right++ {
		tmp[s2[right]-97]--
		for tmp[s2[right]-97] < 0 {
			tmp[s2[left]-97]++
			left++
		}
		if sumZero(tmp) {
			return true
		}
	}
	return false
}
