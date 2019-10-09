package integers

import "testing"

func TestAdd(t *testing.T) {
	t.Run("1 plus 3", func(t *testing.T) {
		sum := Add(1,3)
		expected := 4

		if sum != expected {
			t.Errorf("expected '%d' but got '%d'", expected, sum)
		}
	})
}
