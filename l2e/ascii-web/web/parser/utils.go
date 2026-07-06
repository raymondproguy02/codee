package parser

import (
	"fmt"
	"strings"
)

func Val(input string) (rune, error) {
	for _, r := range input {
		if r != '\r' && r != '\n' && (r < ' ' || r > '~') {
			return r, fmt.Errorf("invalid character")
		}
	}
	return 0, nil
}

func Split(input string) []string {
	input = strings.ReplaceAll(input, "\r\n", "\n")
	return strings.Split(input, "\n")
}
