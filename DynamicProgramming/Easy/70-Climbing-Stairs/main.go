package main

func climbStairs(n int) int {
	if n < 2 {
		return 1
	}
	first, second := 1, 2
	for i := 2; i < n; i++ {
		first, second = second, first+second
	}
	return second
}
