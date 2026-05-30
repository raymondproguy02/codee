package main

import (
	"strings"
)

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
	if style != "normal" && style != "bold" &&
		style != "italic" && style != "outline" {
		panic("invalid style")
	}

	a.style = style
	return a
}

func (a *ArtBuilder) Build() string {
	lines := make([]string, 8)

	for i := 0; i < 8; i++ {

		switch a.style {

		case "bold":
			for _, r := range a.text {
				lines[i] += string(r) + string(r)
			}

		case "italic":
			lines[i] = strings.Repeat(" ", 7-i) + a.text

		case "outline":
			border := strings.Repeat("*", len(a.text)+4)

			if i == 0 || i == 7 {
				lines[i] = border
			} else {
				lines[i] = "* " + a.text + " *"
			}

		default:
			lines[i] = a.text
		}
	}

	return strings.Join(lines, "\n")
}
