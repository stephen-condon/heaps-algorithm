package heaps

import (
	"reflect"
	"testing"
)

func TestCanary(t *testing.T) {

}

func TestPermutationsInt(t *testing.T) {
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

func TestPermutationsString(t *testing.T) {
	want3 := [][]string{{"A", "B", "C"}, {"B", "A", "C"}, {"C", "A", "B"}, {"A", "C", "B"}, {"B", "C", "A"}, {"C", "B", "A"}}
	want4 := [][]string{{"A", "B", "C", "D"}, {"B", "A", "C", "D"}, {"C", "A", "B", "D"}, {"A", "C", "B", "D"}, {"B", "C", "A", "D"}, {"C", "B", "A", "D"}, {"D", "B", "C", "A"}, {"B", "D", "C", "A"}, {"C", "D", "B", "A"}, {"D", "C", "B", "A"}, {"B", "C", "D", "A"}, {"C", "B", "D", "A"}, {"D", "A", "C", "B"}, {"A", "D", "C", "B"}, {"C", "D", "A", "B"}, {"D", "C", "A", "B"}, {"A", "C", "D", "B"}, {"C", "A", "D", "B"}, {"D", "A", "B", "C"}, {"A", "D", "B", "C"}, {"B", "D", "A", "C"}, {"D", "B", "A", "C"}, {"A", "B", "D", "C"}, {"B", "A", "D", "C"}}
	type test struct {
		slice []string
		want  [][]string
	}

	tests := []test{
		{[]string{"A", "B", "C"}, want3},
		{[]string{"A", "B", "C", "D"}, want4},
	}

	for _, tc := range tests {
		got := Permutations(tc.slice)
		if !reflect.DeepEqual(tc.want, got) {
			t.Error("Incorrect calculation of permutations")
		}
	}
}

func TestPermutationsFloat64(t *testing.T) {
	want3 := [][]float64{{1.23, 2.34, 3.45}, {2.34, 1.23, 3.45}, {3.45, 1.23, 2.34}, {1.23, 3.45, 2.34}, {2.34, 3.45, 1.23}, {3.45, 2.34, 1.23}}
	want4 := [][]float64{{1.23, 2.34, 3.45, 4.56}, {2.34, 1.23, 3.45, 4.56}, {3.45, 1.23, 2.34, 4.56}, {1.23, 3.45, 2.34, 4.56}, {2.34, 3.45, 1.23, 4.56}, {3.45, 2.34, 1.23, 4.56}, {4.56, 2.34, 3.45, 1.23}, {2.34, 4.56, 3.45, 1.23}, {3.45, 4.56, 2.34, 1.23}, {4.56, 3.45, 2.34, 1.23}, {2.34, 3.45, 4.56, 1.23}, {3.45, 2.34, 4.56, 1.23}, {4.56, 1.23, 3.45, 2.34}, {1.23, 4.56, 3.45, 2.34}, {3.45, 4.56, 1.23, 2.34}, {4.56, 3.45, 1.23, 2.34}, {1.23, 3.45, 4.56, 2.34}, {3.45, 1.23, 4.56, 2.34}, {4.56, 1.23, 2.34, 3.45}, {1.23, 4.56, 2.34, 3.45}, {2.34, 4.56, 1.23, 3.45}, {4.56, 2.34, 1.23, 3.45}, {1.23, 2.34, 4.56, 3.45}, {2.34, 1.23, 4.56, 3.45}}
	type test struct {
		slice []float64
		want  [][]float64
	}

	tests := []test{
		{[]float64{1.23, 2.34, 3.45}, want3},
		{[]float64{1.23, 2.34, 3.45, 4.56}, want4},
	}

	for _, tc := range tests {
		got := Permutations(tc.slice)
		if !reflect.DeepEqual(tc.want, got) {
			t.Error("Incorrect calculation of permutations")
		}
	}
}
