package year2020

import "testing"

func TestGetTotalSquareFeetNeeded(t *testing.T) {
	tests := []struct {
		name string
		args Dimension
		want int
	}{
		{"Advent Test 1", Dimension{2, 3, 4}, 58},
		{"Advent Test 2", Dimension{1, 1, 10}, 43},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getTotalSquareFeetNeeded(tt.args); got != tt.want {
				t.Errorf("getTotalSquareFeetNeeded() = %v, want %v", got, tt.want)
			}
		})
	}
}
