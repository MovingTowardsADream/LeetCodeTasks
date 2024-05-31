package main

func lengthOfLastWord(s string) int {
	sum, leight := 0, len(s)-1
	for leight >= 0 && s[leight] == ' ' {
		leight--
	}
	for i := leight; i >= 0; i-- {
		if s[i] == ' ' {
			break
		}
		sum++
	}
	return sum
}
