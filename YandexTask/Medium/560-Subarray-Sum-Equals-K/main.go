package main

func subarraySum(nums []int, k int) int {
	tmp := make(map[int][]int, 0)
	sum, amount := 0, 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if sum == k {
			amount++
		}
		amount += len(tmp[sum-k])
		tmp[sum] = append(tmp[sum], i)
	}
	return amount
}
