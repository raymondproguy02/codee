package perser

import (
	"testing"
)

func TestGenerateArt(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		fonts []string
		words []string
		want  string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GenerateArt(tt.fonts, tt.words)
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("GenerateArt() = %v, want %v", got, tt.want)
			}
		})
	}
}
