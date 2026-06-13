package perser

import (
	"testing"
)

func TestSplitInput(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		input string
		want  []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SplitInput(tt.input)
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("SplitInput() = %v, want %v", got, tt.want)
			}
		})
	}
}
