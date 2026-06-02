package main

func g(c rune, r int) (s string) {
	for i := 0; i < 8; i++ {
		if (int(c)^(r*17)^(i*11))%3 != 0 {
			s += "*"
		} else {
			s += " "
		}
	}
	return
}

func GenerateFont() map[rune][]string {
	f := map[rune][]string{' ': make([]string, 8)}
	for i := range f[' '] {
		f[' '][i] = "        "
	}
	for c := rune(33); c <= 126; c++ {
		f[c] = []string{g(c, 0), g(c, 1), g(c, 2), g(c, 3), g(c, 4), g(c, 5), g(c, 6), g(c, 7)}
	}
	return f
}
