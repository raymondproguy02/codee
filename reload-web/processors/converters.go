package processors

import (
	"strconv"
	"strings"
)

func Hex(text string, base int64) string {
	word := strings.Fields(text)
	for i := 0; i < len(word); i++ {
		if word[i] == "(hex)" {
			conv, err := strconv.ParseInt(word[i-1], 16, 64)
			if err != nil {
				return "Error: Invalid hexadecimal value"
			}
			word[i-1] = strconv.FormatInt(conv, 10)
			word = append(word[:i], word[i+1:]...)
		}
	}
	return strings.Join(word, " ")
}

func Bin(text string, base int64) string {
	word := strings.Fields(text)
	for i := 0; i < len(word); i++ {
		if word[i] == "(bin)" {
			conv, err := strconv.ParseInt(word[i-1], 16, 64)
			if err != nil {
				return "Error: Invalid binary value"
			}
			word[i-1] = strconv.FormatInt(conv, 10)
			word = append(word[:i], word[i+1:]...)
		}
	}
	return strings.Join(word, " ")
}
