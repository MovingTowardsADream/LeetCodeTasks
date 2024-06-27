package main

import (
	"fmt"
)

func compress(chars []byte) int {
	if len(chars) < 2 {
		return len(chars)
	}
	first, local_amount := 0, 0
	ch := chars[0]
	for right := 0; right < len(chars); right++ {
		if ch == chars[right] {
			local_amount++
		} else {
			first++
			if local_amount != 1 {
				str := []byte(fmt.Sprintf("%d", local_amount))
				local_amount = 1
				copy(chars[first:], str)
				first += len(str)
			}
			ch = chars[right]
			chars[first] = chars[right]
		}
	}
	first++
	if local_amount != 1 {
		str := []byte(fmt.Sprintf("%d", local_amount))
		copy(chars[first:], str)
		first += len(str)
	}
	return first
}
