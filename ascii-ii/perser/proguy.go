package perser

import (
	"fmt"
	"os"
	"strings"
)

func Generate(font []string, words []string) string {
	var res strings.Builder
	for index, word := range words {
		if word == "" {
			res.WriteString("\n")
			continue
		}

		for i := 1; i <= 8; i++ {
			for _, ch := range word {
				res.WriteString(font[i+(int(ch-32)*9)])
			}
			res.WriteString("\n")
		}

		if index < len(words)-1 && words[index+1] != "" {
			res.WriteString("\n")
		}
	}
	return res.String()
}

func Load(file string) ([]string, error) {
	data, err := os.ReadFile("banners/" + file + ".txt")
	if err != nil {
		return nil, err
	}
	var line = strings.Split(strings.ReplaceAll(string(data), "\r\n", "\n"), "\n")
	return line, nil
}

func Split(input string) []string {
	input = strings.ReplaceAll(input, "\r\n", "\n")
	return strings.Split(input, "\\n")
}

func Validate(input string) (rune, error) {
	for _, r := range input {
		if r != '\r' && r != '\n' && (r < ' ' || r > '~') {
			return r, fmt.Errorf("invalid character")
		}
	}
	return 0, nil
}
