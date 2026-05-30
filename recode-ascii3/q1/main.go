package main

func GeneratePattern(c rune) []string {
	pattern := map[rune][]string{}
	for i := 'A'; i < 'Z'; i++ {
		pattern[i] = []string{
			"  ##  ",
			" #  # ",
			" #  # ",
			" #### ",
			" #  # ",
			" #  # ",
			" #  # ",
			"      ",
		}
		pattern['Z'] = []string{
			" #### ",
			"    # ",
			"   #  ",
			"  #   ",
			" #    ",
			" #    ",
			" #### ",
			"      ",
		}
	}
	if lines, ok := pattern[c]; ok {
		return lines
	}

	return nil
}
