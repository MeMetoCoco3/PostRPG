package main

import "testing"

func TestDistanceBetweenTwoPoints(t *testing.T) {

	tests := []struct {
		x1, y1, x2, y2 int
		expected       int
	}{
		{1, 5, 2, 4, 2},
		{0, 0, 3, 0, 3},
		{0, 0, 0, 4, 4},
		{1, 1, 1, 1, 0},
		{2, 3, 5, 7, 7},
	}

	for _, tt := range tests {
		result := DistanceBetweenTwoPoints(tt.x1, tt.y1, tt.x2, tt.y2)
		if result != tt.expected {
			t.Errorf("DBTP(%d, %d, %d, %d) = %d; want %d", tt.x1, tt.y1, tt.x2, tt.y2, result, tt.expected)
		}
	}
}
