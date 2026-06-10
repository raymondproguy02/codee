package main

import (
	"fmt"
	"os"
	"strings"
)

func GenerateArt(fontFileLines []string, words []string) string {
	var result strings.Builder
	for index, word := range words {
		if word == "" {
			continue
		}
		for i := 1; i < 9; i++ {
			for _, char := range word {
				result.WriteString(fontFileLines[i+(int(char-32)*9)])
			}
			if i < 8 {
				result.WriteString("\n")
			}
		}
		if index < len(words)-1 {
			result.WriteString("\n")
		}
	}
	return result.String()
}

func LoadBanner(file string) ([]string, error) {
	data, err := os.ReadFile("banners/" + file + ".txt")
	if err != nil {
		return nil, err
	}
	var lines = strings.Split(string(data), "\n")
	if len(lines) == 0 {
		return nil, fmt.Errorf("empty file")
	}
	return lines, nil
}

func Split(input string) []string {
	input = strings.ReplaceAll(input, "\r\n", "\n")
	return strings.Split(input, "\n")
}

func ValidateInput(input string) (rune, error) {
	for _, r := range input {
		if r != '\n' && r != '\r' && (r < ' ' || r > '~') {
			return r, fmt.Errorf("unsupported character: %q", r)
		}
	}
	return 0, nil
}

func main() {
	if len(os.Args) == 3 {
		if len(os.Args[1]) == 0 {
			return
		}
		fontFileLines, err := LoadBanner(os.Args[2])
		if err != nil {
			fmt.Printf("Error: %v", err)
			return
		}
		words := Split(os.Args[1])
		result := GenerateArt(fontFileLines, words)
		fmt.Println(result)
	}
}
