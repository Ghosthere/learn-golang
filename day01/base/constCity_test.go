package base

import "testing"

func TestCity(t *testing.T) {
	tests := []struct {
		name string
		want uint
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := City(); got != tt.want {
				t.Errorf("City() = %v, want %v", got, tt.want)
			}
		})
	}
}
