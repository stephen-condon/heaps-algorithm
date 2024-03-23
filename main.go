package main

import (
	"fmt"
	"permutations/heaps"
)

func main() {
	test3 := []int{1, 2, 3}
	result := heaps.Permutations(test3)
	// should be [[1 2 3] [2 1 3] [3 1 2] [1 3 2] [2 3 1] [3 2 1]]

	test4 := []int{1, 2, 3, 4}
	result2 := heaps.Permutations(test4)
	// should be [[1 2 3 4] [2 1 3 4] [3 1 2 4] [1 3 2 4] [2 3 1 4] [3 2 1 4] [4 2 3 1] [2 4 3 1] [3 4 2 1] [4 3 2 1] [2 3 4 1] [3 2 4 1] [4 1 3 2] [1 4 3 2] [3 4 1 2] [4 3 1 2] [1 3 4 2] [3 1 4 2] [4 1 2 3] [1 4 2 3] [2 4 1 3] [4 2 1 3] [1 2 4 3] [2 1 4 3]]

	fmt.Println(result)
	fmt.Println(result2)
}
