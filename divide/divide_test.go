package divide

import "testing"

func TestDivide(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		got, err := Divide(10, 2)
		if err != nil {
			t.Fatalf("Divide(10, 2) returned unexpected error: %v", err)
		}
		if got != 5 {
			t.Errorf("Divide(10, 2) = %d; want 5", got)
		}
	})

	t.Run("negative dividend", func(t *testing.T) {
		got, err := Divide(-9, 3)
		if err != nil {
			t.Fatalf("Divide(-9, 3) returned unexpected error: %v", err)
		}
		if got != -3 {
			t.Errorf("Divide(-9, 3) = %d; want -3", got)
		}
	})

	t.Run("division by zero", func(t *testing.T) {
		_, err := Divide(1, 0)
		if err == nil {
			t.Errorf("Divide(1, 0) returned nil error; want non-nil")
		}
	})
}
