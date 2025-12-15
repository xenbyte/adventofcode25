package util

import "strings"

func ParseInput(input string) []string {
	var chars []string
	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		chars = append(chars, line)
	}
	return chars
}
