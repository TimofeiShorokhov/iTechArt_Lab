package tasks

import "strings"

func ReverseParentheses(s string) string {
	stack := make([]string, 0)

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, string(s[i]))
		} else if s[i] == ')' {
			temp := make([]string, 0)
			for stack[len(stack)-1] != string('(') {
				temp = append(temp, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = stack[:len(stack)-1]
			for i := 0; i < len(temp); i++ {
				stack = append(stack, temp[i])
			}
		} else {
			stack = append(stack, string(s[i]))
		}
	}
	return strings.Join(stack, "")
}
