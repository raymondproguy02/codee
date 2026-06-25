package main

import (
	"fmt"
	"strconv"
	"strings"
)

func up(s string) string {
	words := strings.Fields(s)

	for i := 0; i < len(words); i++ {
		if words[i] == "(up)" {
			preW := strings.ToUpper(words[i-1])
			words[i-1] = preW

			words = append(words[:i], words[i+1:]...)

		}
	}
	return strings.Join(words, " ")
}

func low(s string) string {
	words := strings.Fields(s)

	for i := 0; i < len(words); i++ {
		if words[i] == "(low)" {
			preW := strings.ToLower(words[i-1])
			words[i-1] = preW

			words = append(words[:i], words[i+1:]...)

		}
	}
	return strings.Join(words, " ")
}

func cap(s string) string {
	words := strings.Fields(s)

	for i := 0; i < len(words); i++ {
		if words[i] == "(cap)" {
			preW := strings.ToUpper(words[i-1][:1]) + strings.ToLower(words[i-1][1:])
			words[i-1] = preW

			words = append(words[:i], words[i+1:]...)

		}
	}
	return strings.Join(words, " ")
}

func upN(s string) string {
	words := strings.Fields(s)

	for i := 0; i < len(words); i++ {
		if words[i] == "(up," && strings.HasSuffix(words[i+1], ")") && i+1 <= len(words) {
			words[i+1] = strings.TrimSuffix(words[i+1], ")")
			count, err := strconv.Atoi(words[i+1])
			if err != nil {
				continue
			}
			start := i - count
			if start < 0 {
				start = 0
			}

			for j := start; j < i; j++ {
				words[j] = strings.ToUpper(words[j])
			}
			words = append(words[:i], words[i+2:]...)
			i--

		} else {
			continue
		}
	}
	return strings.Join(words, " ")
}

func lowN(s string) string {
	words := strings.Fields(s)

	for i := 0; i < len(words); i++ {
		if words[i] == "(low," && strings.HasSuffix(words[i+1], ")") && i+1 <= len(words) {
			words[i+1] = strings.TrimSuffix(words[i+1], ")")
			count, err := strconv.Atoi(words[i+1])
			if err != nil {
				continue
			}
			start := i - count
			if start < 0 {
				start = 0
			}

			for j := start; j < i; j++ {
				words[j] = strings.ToLower(words[j])
			}
			words = append(words[:i], words[i+2:]...)
			i--

		} else {
			continue
		}
	}
	return strings.Join(words, " ")
}

func capN(s string) string {
	words := strings.Fields(s)

	for i := 0; i < len(words); i++ {
		if words[i] == "(cap," && strings.HasSuffix(words[i+1], ")") && i+1 <= len(words) {
			words[i+1] = strings.TrimSuffix(words[i+1], ")")
			count, err := strconv.Atoi(words[i+1])
			if err != nil {
				continue
			}
			start := i - count
			if start < 0 {
				start = 0
			}

			for j := start; j < i; j++ {
				words[j] = strings.ToUpper(words[j][:1])+strings.ToLower(words[j][1:])
			}
			words = append(words[:i], words[i+2:]...)
			i--

		} else {
			continue
		}
	}
	return strings.Join(words, " ")
}


