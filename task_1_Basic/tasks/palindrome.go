package tasks

import (
	"strings"
)

func IsPalindrome(str string) bool {
	u := strings.ToLower(str)
	r := []rune(u)
	var res []rune
	for i := len(r) - 1; i >= 0; i-- {
		res = append(res, r[i])
	}
	strings.ToLower(string(res))
	if string(res) == u {
		return true
	}
	return false
}
