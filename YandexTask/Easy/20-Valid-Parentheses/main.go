package main

func isValid(s string) bool {
	stack := make([]rune, 0)
	leight := 0
	for _, i := range s {
		if len(stack) == 0 || i == '(' || i == '{' || i == '[' {
			stack = append(stack, i)
			leight++
		} else if leight == 0 {
			return false
		} else if (i == ')' && stack[leight-1] == '(') || (i == '}' && stack[leight-1] == '{') || (i == ']' && stack[leight-1] == '[') {
			leight--
			stack = stack[:leight]
		} else {
			return false
		}
	}
	if len(stack) == 0 {
		return true
	}
	return false
}
