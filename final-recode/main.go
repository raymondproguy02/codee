package main

func main() {
	func LoadBanner(file string) ([]string, error) {
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

func SplitInput(input string) []string {
	input = strings.ReplaceAll(input, "\r\n", "\n")
	return strings.Split(input, "\n")
}

func Validate(input string) (rune, error) {
	for _, r := range input {
		if r != '\r' && r != '\n' && (r < ' ' || r > '~') {
			return r, fmt.Errorf("Invalid character")
		}
	}
	return 0, nil
}

func GenerateArt(fonts []string, words []string) string {
	var res strings.Builder
	for index, word := range words {
		if word == "" {
			res.WriteString("\n")
			continue
		}
		for i := 1; i <= 8; i++ {
			for _, ch := range word {
				res.WriteString(fonts[i+(int(ch-32)*9)])
			}
			res.WriteString("\n")
		}
		if index < len(words)-1 && words[index+1] != "" {
			res.WriteString("\n")
		}
	}
	return res.String()
}
}