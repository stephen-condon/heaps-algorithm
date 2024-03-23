package heaps

// maybe try generics?
func Permutations(slice []int) [][]int {
	result := [][]int{}
	var generate func(int, []int)

	generate = func(k int, slice []int) {
		if k == 1 {
			tmp := make([]int, len(slice))
			copy(tmp, slice)
			result = append(result, tmp)
		} else {
			for i := 0; i < k; i++ {
				generate(k-1, slice)
				if k%2 == 0 {
					swap(i, k-1, slice)
				} else {
					swap(0, k-1, slice)
				}
			}
		}
	}

	generate(len(slice), slice)

	return result
}

func swap(n, m int, s []int) {
	el := s[n]
	s[n] = s[m]
	s[m] = el
}
