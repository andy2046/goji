package intutil

import (
	"testing"
)

func TestUniq(t *testing.T) {
	tests := []struct {
		in  []int
		out []int
	}{
		{[]int{5, 4, 3, 4, 5}, []int{5, 4, 3}},
		{[]int{4, 4, 2, 2, 4}, []int{4, 2}},
	}

	for _, test := range tests {
		t.Run("Uniq", func(t *testing.T) {
			i := NewInt()
			r := i.Uniq(test.in)

			if !i.Equal(r, test.out) {
				t.Errorf("got %v, want %v", r, test.out)
			}
		})
	}

}
