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

func LoadBanner(fileName string) ([]string, error) {
	fontFile, err := os.ReadFile("banner/" + fileName + ".txt")
	if err != nil {
		return nil, err
	}
	var fontFileLines = strings.Split(string(fontFile), "\n")
	return fontFileLines, nil
}

func SplitInput(input string) []string {
	input = strings.ReplaceAll(input, "\r\n", "\n")
	return strings.Split(input, "\n")
}

func ValidateInput(input string) ([]rune, error) {
	odds := []rune{}
	for _, char := range input {
		if char != '\n' && char != '\r' && (char < ' ' || char > '~') {
			odds = append(odds, char)
			continue
		}
	}
	if len(odds) > 0 {
		return odds, fmt.Errorf("unsupported character: %c", odds)
	}
	return nil, nil
}

func main() {
	if len(os.Args) != 3 {
		return
	}
	if len(os.Args[1]) == 0 {
		return
	}
	val, err := ValidateInput(os.Args[1])
	if err != nil {
		fmt.Println("invalid", val)
		return
	}

	font, err := LoadBanner(os.Args[2])
	if err != nil {
		fmt.Println(err)
		return
	}

	words := SplitInput(os.Args[1])
	res := GenerateArt(font, words)
	fmt.Println(res)
}
