package heaps

func Permutations[E any](slice []E) [][]E {
	result := [][]E{}
	var generate func(int, []E)

	generate = func(k int, slice []E) {
		if k == 1 {
			tmp := make([]E, len(slice))
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

func swap[E any](n, m int, s []E) {
	el := s[n]
	s[n] = s[m]
	s[m] = el
}
