package perser

import (
	"fmt"
	"os"
	"strings"
)

// LoadBanner reads and validates the ASCII banner files
func LoadBanner(file string) ([]string, error) {
	// Remove extra extentions if provided by the user
	file = strings.TrimSuffix(file, ".txt")
	data, err := os.ReadFile("banners/" + file + ".txt")
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, fmt.Errorf("Empty banner file")
	}
	var lines = strings.Split(strings.ReplaceAll(string(data), "\r\n", "\n"), "\n")
	if len(lines) < 855 || len(lines) > 856 {
		return nil, fmt.Errorf("Incomplete banner file")
	}
	return lines, nil
}
