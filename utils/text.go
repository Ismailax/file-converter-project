package utils

import (
	"strings"
)

func CleanText(s string) string {
	return strings.ReplaceAll(strings.TrimSpace(s), "\t", "")
}

func Atoi(s string) int {
	n := 0
	for _, c := range s {
		if c >= '0' && c <= '9' {
			n = n*10 + int(c-'0')
		}
	}
	return n
}
