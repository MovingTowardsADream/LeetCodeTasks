package main

func mySqrt(x int) int {
	left, right := 0, x
	for left < right {
		mid := (left + right) / 2
		pow_mid := mid * mid
		if pow_mid > x {
			right = mid - 1
		} else if pow_mid < x {
			left = mid + 1
		} else {
			return mid
		}
	}
	if left*left > x {
		return left - 1
	}
	return left
}
