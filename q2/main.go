package main

import (
	"strings"
)

func pattern(text rune) []string {
	switch text {
	case '0':
		return []string{
			"|   |",
			"|   |",
			"|   |",
			"|   |",
			"|___|",
		}
	case '1':
		return []string{
			"  |  ",
			"  |  ",
			"  |  ",
			"  |  ",
			"  |  ",
		}
	case '2':
		return []string{
			" ___ ",
			"    |",
			" ___|",
			"|    ",
			"|___ ",
		}
	case '3':
		return []string{
			" ___ ",
			"    |",
			" ___|",
			"    |",
			" ___|",
		}
	case '4':
		return []string{
			"|   |",
			"|   |",
			"|___|",
			"    |",
			"    |",
		}
	case '5':
		return []string{
			" ___ ",
			"|    ",
			"|___ ",
			"    |",
			" ___|",
		}
	case '6':
		return []string{
			" ___ ",
			"|    ",
			"|___ ",
			"|   |",
			"|___|",
		}
	case '7':
		return []string{
			" ___ ",
			"    /",
			"   / ",
			"  /  ",
			" /   ",
		}
	case '8':
		return []string{
			" ___ ",
			"|   |",
			"|___|",
			"|   |",
			"|___|",
		}
	case '9':
		return []string{
			" ___ ",
			"|   |",
			"|___|",
			"    |",
			" ___|",
		}
	}
	return []string{}
}
func StringToArt(input string) string {
	if input == "" {
		return ""
	}

	lines := strings.Split(input, "\n")

	var holder []string
	for _, char := range lines {
		rows := make([]string, 5)
		for _, k := range char {
			if k < '0' || k > '9' {
				return ""
			}
			word := pattern(k)
			for i := 0; i < 5; i++ {
				rows[i] += word[i]
			}
		}
		Allchar := strings.Join(rows, "\n")
		holder = append(holder, Allchar)
	}
	return strings.Join(holder, "\n") + "\n"
}
