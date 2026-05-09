package multiply

import "testing"

func TestMultiply(t *testing.T) {
	cases := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"positive", 3, 4, 12},
		{"negatives", -5, -6, 30},
		{"mixed", -7, 8, -56},
		{"with zero", 0, 9, 0},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := Multiply(tc.a, tc.b)
			if got != tc.expected {
				t.Errorf("Multiply(%d, %d) = %d; want %d", tc.a, tc.b, got, tc.expected)
			}
		})
	}
}
