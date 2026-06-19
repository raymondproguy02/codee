package main

import (
	"os"
	"strings"
)

func load(file string) (map[rune][]string, error) {
	data, err := os.ReadFile("banners/" + file + ".txt")
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(data), "\n")
	lines = lines[1:]

	ascii := make(map[rune][]string)
	for i := ' '; i <= '~'; i++ {
		s := int(i-32) * 9
		e := s + 8
		ascii[i] = lines[s:e]
	}
	return ascii, nil

}
