package main

func romanToInt(s string) int {
	valuation := map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	sum, leight := 0, len(s)
	for i := 0; i < leight; i++ {
		if i != leight-1 && valuation[s[i]] < valuation[s[i+1]] {
			sum -= valuation[s[i]]
		} else {
			sum += valuation[s[i]]
		}
	}
	return sum
}
