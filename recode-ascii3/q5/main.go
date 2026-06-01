package main

import (
	"fmt"
	"strings"
)

type Animation struct {
	text   string
	frames int
	data   []string
}

func NewAnimation(text string, frames int) *Animation {
	return &Animation{
		text:   text,
		frames: frames,
	}
}

func (a *Animation) GenerateSpinFrames() {
	for i := 0; i < a.frames; i++ {
		a.data = append(a.data, strings.TrimRight(strings.Repeat(fmt.Sprintf("%s%d\n", a.text, i), 10), "\n"))
	}
}

func (a *Animation) GenerateWaveFrames() {
	a.GenerateSpinFrames()
}

func (a *Animation) GenerateZoomFrames() {
	a.GenerateSpinFrames()
}

func (a *Animation) GetFrame(i int) string {
	if len(a.data) == 0 {
		return ""
	}
	return a.data[i%len(a.data)]
}

func (a *Animation) Play() string {
	return strings.Join(a.data, "\n\n")
}
