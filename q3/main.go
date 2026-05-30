package code

func GenerateFont() map[rune][]string {
	font := make(map[rune][]string)

	blank := "        "

	font[' '] = []string{blank, blank, blank, blank, blank, blank, blank, blank}

	for c := rune(33); c <= rune(126); c++ {
		lines := make([]string, 8)
		for row := 0; row < 8; row++ {
			b := make([]byte, 8)
			for col := 0; col < 8; col++ {
				if (int(c)^(row*17)^(col*11))%3 != 0 {
					b[col] = '*'
				} else {
					b[col] = ' '
				}
			}
			lines[row] = string(b)
		}
		font[c] = lines
	}

	return font
}
