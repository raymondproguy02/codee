package perser

import (
	"fmt"
	"strings"
)

// Split user input into separate words based on newline characters.
func SplitInput(input string) []string {
	input = strings.ReplaceAll(input, "\r\n", "\n")
	return strings.Split(input, "\n")
}

// Validate ensures input contains only printable ASCII characters.
func Validate(input string) (rune, error) {
	for _, r := range input {
		if r != '\r' && r != '\n' && (r < ' ' || r > '~') {
			return r, fmt.Errorf("Invalid character")
		}
	}
	return 0, nil
}
