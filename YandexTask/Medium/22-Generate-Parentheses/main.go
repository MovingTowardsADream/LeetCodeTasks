package main

func permutation(result *[]string, stack []byte, n int, open, close int) {
	if len(stack) == n*2 && open == close {
		tmp := make([]byte, n*2)
		copy(tmp, stack)
		*result = append(*result, string(tmp))
	} else {
		if open < n {
			stack = append(stack, '(')
			permutation(result, stack, n, open+1, close)
			stack = stack[:len(stack)-1]
		}
		if open > close {
			stack = append(stack, ')')
			permutation(result, stack, n, open, close+1)
			stack = stack[:len(stack)-1]
		}
	}
}

func generateParenthesis(n int) []string {
	var result []string
	stack := make([]byte, 0, n*2)
	permutation(&result, stack, n, 0, 0)
	return result
}

// or
//
//func generateParenthesis(n int) []string {
//	result := make([]string, 0)
//	stack := make([]byte, 0, n*2)
//
//	var permutation func(int, int)
//	permutation = func(open, close int) {
//		if open + close == n * 2 {
//			result = append(result, string(stack))
//			return
//		}
//		if open < n {
//			stack = append(stack, '(')
//			permutation(open+1, close)
//			stack = stack[:len(stack)-1]
//		}
//		if close < open {
//			stack = append(stack, ')')
//			permutation(open, close+1)
//			stack = stack[:len(stack)-1]
//		}
//	}
//	permutation(0, 0)
//	return result
//}
