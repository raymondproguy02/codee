package main

import "strings"

type ArtBuilder struct {
	text  string
	style string
}

func NewArtBuilder() *ArtBuilder {
	return &ArtBuilder{style: "normal"}
}

func (a *ArtBuilder) AddText(text string) *ArtBuilder {
	a.text += text
	return a
}

func (a *ArtBuilder) SetStyle(style string) *ArtBuilder {
	if style != "normal" && style != "bold" && style != "italic" && style != "outline" {
		panic("invalid style")
	}
	a.style = style
	return a
}

func (a *ArtBuilder) Build() string {
	lines := make([]string, 8)
	w := len(a.text)

	for i := range lines {
		switch a.style {
		case "bold":
			for _, r := range a.text {
				lines[i] += string(r) + string(r)
			}
		case "italic":
			pad := strings.Repeat(" ", (7 - i))
			lines[i] = pad + a.text + strings.Repeat(" ", i)
		case "outline":
			bar := "+" + strings.Repeat("-", w+2) + "+"
			if i == 0 || i == 7 {
				lines[i] = bar
			} else {
				lines[i] = "| " + a.text + " |"
			}
		default:
			lines[i] = a.text
		}
	}

	return strings.Join(lines, "\n") + "\n"
}