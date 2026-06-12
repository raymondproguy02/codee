package perser

import (
	"testing"
)

func TestLoadBanner(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		file    string
		want    []string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := LoadBanner(tt.file)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("LoadBanner() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("LoadBanner() succeeded unexpectedly")
			}
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("LoadBanner() = %v, want %v", got, tt.want)
			}
		})
	}
}
