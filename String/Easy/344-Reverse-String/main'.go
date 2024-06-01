package main

func reverseString(s []byte) {
	leight := len(s)
	for i := 0; i < leight/2; i++ {
		s[i], s[leight-i-1] = s[leight-i-1], s[i]
	}
}
