package parser

import "strings"

func GenArt(font []string, words []string) string {
	var res strings.Builder
	for index, word := range words {
		if word == "" {
			res.WriteString("\n")
			continue
		}
		for i := 1; i <= 8; i++ {
			for _, ch := range word {
				res.WriteString(font[i+int(ch-32)*9])
			}
			res.WriteString("\n")
		}
		if index < len(words)-1 && words[index+1] != "" {
			res.WriteString("\n")
		}
	}
	return res.String()
}
