package parser

import (
	"fmt"
	"os"
	"strings"
)

func LoadBanner(file string) ([]string, error) {
	file = strings.TrimSuffix(file, ".txt")
	data, err := os.ReadFile("banners/" + file + ".txt")

	if len(data) == 0 {
		return nil, fmt.Errorf("empty file")
	}
	if err != nil {
		return nil, err
	}

	lines := strings.Split(strings.ReplaceAll(string(data), "\r\n", "\n"), "\n")
	if len(lines) < 855 || len(lines) > 856 {
		return nil, fmt.Errorf("incomplete file")
	}
	return lines, nil
}
