package main

import "strconv"

func compress(chars []byte) int {
	left := 0
	count := 0
	if len(chars) == 0 {
		return 0
	}
	for right := 0; right < len(chars); right++ {
		if chars[left] == chars[right] {
			count++
		} else {
			left++
			if count != 1 {
				tmp := strconv.Itoa(count)
				for i := 0; i < len(tmp); i++ {
					chars[left+i] = tmp[i]
				}
				left += len(tmp)
			}
			count = 1
			chars[left] = chars[right]
		}
	}
	left++
	if count != 1 {
		tmp := strconv.Itoa(count)
		for i := 0; i < len(tmp); i++ {
			chars[left+i] = tmp[i]
		}
		left += len(tmp)
	}
	return left
}
