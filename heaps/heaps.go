package heaps

// maybe try generics?
func Permutations(slice baseSlice) [][]int {
	result := [][]int{}
	var generate func(int, baseSlice)

	generate = func(k int, slice baseSlice) {
		if k == 1 {
			tmp := make(baseSlice, len(slice))
			copy(tmp, slice)
			result = append(result, tmp)
		} else {
			for i := 0; i < k; i++ {
				generate(k-1, slice)
				if k%2 == 0 {
					slice.swap(i, k-1)
				} else {
					slice.swap(0, k-1)
				}
			}
		}
	}

	generate(len(slice), slice)

	return result
}

func (bs baseSlice) swap(n, m int) {
	el := bs[n]
	bs[n] = bs[m]
	bs[m] = el
}

type baseSlice []int
