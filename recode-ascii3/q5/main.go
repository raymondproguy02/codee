package main

import (
	"strconv"
	"strings"
)

type Animation struct {
	text string
	n    int
	data []string
}

func NewAnimation(t string, n int) *Animation {
	return &Animation{t, n, nil}
}
func (a *Animation) GenerateSpinFrames() { a.make(0) }
func (a *Animation) GenerateWaveFrames() { a.make(1) }
func (a *Animation) GenerateZoomFrames() { a.make(2) }

func (a *Animation) make(mode int) {
	a.data = nil

	for i := 0; i < a.n; i++ {
		f := make([]string, 10)

		for j := range f {
			f[j] = strings.Repeat(" ", 20)
		}

		switch mode {

		case 0:
			f[i%10] = pad(a.text)

		case 1:
			f[3+i%4] = pad(a.text)

		case 2:
			s := ""
			for _, c := range a.text {
				s += strings.Repeat(string(c), i%3+1)
			}
			f[4] = pad(s)
		}

		a.data = append(a.data, strings.Join(f, "\n"))
	}
}

func (a *Animation) GetFrame(i int) string {
	if len(a.data) == 0 {
		return ""
	}
	return a.data[i%len(a.data)]
}

func (a *Animation) Play() string {
	s := ""
	for i, f := range a.data {
		s += "Frame " + strconv.Itoa(i) + "\n" + f + "\n---\n"
	}
	return s
}

func pad(s string) string {
	if len(s) > 20 {
		return s[:20]
	}
	return s + strings.Repeat(" ", 20-len(s))
}
