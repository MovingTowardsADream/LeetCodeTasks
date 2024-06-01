package main

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	tmp_s := make(map[byte]int)
	tmp_t := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		if tmp_s[s[i]] != tmp_t[t[i]] {
			return false
		}
		tmp_s[s[i]], tmp_t[t[i]] = i+1, i+1
	}
	return true
}
