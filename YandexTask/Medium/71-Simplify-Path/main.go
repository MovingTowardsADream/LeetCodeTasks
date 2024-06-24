package main

import "strings"

func simplifyPath(path string) string {
	stack := make([]string, 0)
	names := strings.Split(path, "/")
	for _, name := range names {
		if name == "" || name == "." {
			continue
		} else if name == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, name)
		}
	}
	return "/" + strings.Join(stack, "/")
}
