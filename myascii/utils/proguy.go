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
    nput = strings.ReplaceAll(input, "\r\n", "\n")
    return strings.Split(input, "\n")
}