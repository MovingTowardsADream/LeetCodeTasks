package main

func minRemoveToMakeValid(s string) string {
	stack := make([][]int, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, []int{'(', i})
		} else if s[i] == ')' {
			if len(stack) != 0 && stack[len(stack)-1][0] == '(' {
				stack = stack[:len(stack)-1]
			} else {
				stack = append(stack, []int{')', i})
			}
		}
	}
	result := make([]byte, 0)
	index := 0
	for i := 0; i < len(s); i++ {
		if index < len(stack) && i == stack[index][1] {
			index++
		} else {
			result = append(result, s[i])
		}
	}
	return string(result)
}
