package main

import "testing"

// TestPow tests a Pow(x, y) function, which returns integer x**y.
func TestPow(t *testing.T) {
	for _, tc := range []struct {
		x, y uint
		want uint
	}{
		{x: 0, y: 0, want: 1},
		{x: 1, y: 0, want: 1},
		{x: 7, y: 0, want: 1},
		{x: 0, y: 7, want: 0},
		{x: 0, y: 13467, want: 0},
		{x: 2, y: 3, want: 8},
		{x: 5, y: 7, want: 78125},
		{x: 13, y: 8, want: 815730721},
		{x: 1, y: 1, want: 1},
		{x: 3, y: 3, want: 27},
		{x: 7, y: 7, want: 823543},
	} {
		if got := pow(tc.x, tc.y); got != tc.want {
			t.Errorf("pow(%v, %v) = (%v), want = (%v)", tc.x, tc.y, got, tc.want)
		}
	}
}
