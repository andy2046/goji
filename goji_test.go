package goji

import (
	"testing"
)

func TestFilter(t *testing.T) {
	tests := []struct {
		in     []int
		out    []int
		filter func([]int) func(i int) bool
	}{
		{[]int{5, 4, 3, 2, 1}, []int{5, 4, 3}, func(in []int) func(i int) bool {
			return func(i int) bool {
				if in[i] > 2 {
					return true
				}
				return false
			}
		}},
		{[]int{5, 4, 3, 2, 1}, []int{4, 2}, func(in []int) func(i int) bool {
			return func(i int) bool {
				if in[i]%2 == 0 {
					return true
				}
				return false
			}
		}},
	}

	for _, test := range tests {
		t.Run("Filter", func(t *testing.T) {
			var result []int
			s := Filter(
				len(test.in),
				test.filter(test.in))
			for _, v := range s {
				result = append(result, test.in[v])
			}
			for i, v := range result {
				if v != test.out[i] {
					t.Errorf("got %v, want %v", result, test.out)
				}
			}
		})
	}

}
