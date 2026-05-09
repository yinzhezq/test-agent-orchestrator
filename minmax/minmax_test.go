package minmax

import "testing"

func TestMin(t *testing.T) {
	cases := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"a less than b", 2, 5, 2},
		{"a greater than b", 9, 4, 4},
		{"equal values", 7, 7, 7},
		{"negatives", -3, -8, -8},
		{"mixed sign", -2, 3, -2},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := Min(tc.a, tc.b)
			if got != tc.expected {
				t.Errorf("Min(%d, %d) = %d; want %d", tc.a, tc.b, got, tc.expected)
			}
		})
	}
}

func TestMax(t *testing.T) {
	cases := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"a less than b", 2, 5, 5},
		{"a greater than b", 9, 4, 9},
		{"equal values", 7, 7, 7},
		{"negatives", -3, -8, -3},
		{"mixed sign", -2, 3, 3},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := Max(tc.a, tc.b)
			if got != tc.expected {
				t.Errorf("Max(%d, %d) = %d; want %d", tc.a, tc.b, got, tc.expected)
			}
		})
	}
}
