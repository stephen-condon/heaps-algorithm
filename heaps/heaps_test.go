package heaps

import (
	"reflect"
	"testing"
)

func TestCanary(t *testing.T) {

}

func TestPermutations(t *testing.T) {
	want3 := [][]int{{1, 2, 3}, {2, 1, 3}, {3, 1, 2}, {1, 3, 2}, {2, 3, 1}, {3, 2, 1}}
	want4 := [][]int{{1, 2, 3, 4}, {2, 1, 3, 4}, {3, 1, 2, 4}, {1, 3, 2, 4}, {2, 3, 1, 4}, {3, 2, 1, 4}, {4, 2, 3, 1}, {2, 4, 3, 1}, {3, 4, 2, 1}, {4, 3, 2, 1}, {2, 3, 4, 1}, {3, 2, 4, 1}, {4, 1, 3, 2}, {1, 4, 3, 2}, {3, 4, 1, 2}, {4, 3, 1, 2}, {1, 3, 4, 2}, {3, 1, 4, 2}, {4, 1, 2, 3}, {1, 4, 2, 3}, {2, 4, 1, 3}, {4, 2, 1, 3}, {1, 2, 4, 3}, {2, 1, 4, 3}}
	type test struct {
		slice []int
		want  [][]int
	}

	tests := []test{
		{[]int{1, 2, 3}, want3},
		{[]int{1, 2, 3, 4}, want4},
	}

	for _, tc := range tests {
		got := Permutations(tc.slice)
		if !reflect.DeepEqual(tc.want, got) {
			t.Error("Incorrect calculation of permutations")
		}
	}
}
