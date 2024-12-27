package main

import "testing"

func TestDistanceBetweenTwoPoints(t *testing.T) {

	tests := []struct {
		x1, y1, x2, y2 int
		e1, e2         int
	}{
		{1, 5, 2, 4, 1, 1},
		{0, 0, 3, 0, 3, 0},
		{0, 0, 0, 4, 0, 4},
		{1, 1, 1, 1, 0, 0},
		{2, 3, 5, 7, 3, 4},
	}

	for _, tt := range tests {
		r1, r2 := DistanceBetweenTwoPoints(tt.x1, tt.y1, tt.x2, tt.y2)
		if r1 != tt.e1 || r2 != tt.e2 {
			t.Errorf("DBTP(%d, %d, %d, %d) = (%d, %d); want (%d, %d)", tt.x1, tt.y1, tt.x2, tt.y2, r1, r2, tt.e1, tt.e2)
		}
	}
}
