# Heaps Algorithm

## Background

The Heaps Algorithm is a permutation generation algorithm, outlined [here](https://en.wikipedia.org/wiki/Heap%27s_algorithm).

## Implementation

This implementation uses generics in Go to be able to generate a slice of all possible permutations of the provided slice. Example usage in main.go in this repository:

```
import "https://github.com/stephen-condon/heaps-algorithm/permutations/heaps"

test3 := []int{1, 2, 3}
result := heaps.Permutations(test3)
// should be [[1 2 3] [2 1 3] [3 1 2] [1 3 2] [2 3 1] [3 2 1]]
```

## Run Example Integration

After cloning the repository, run:

```
go run .
```

## Testing

Test the cloned repository with:

```
go test ./...
```

Run the benchmark:

```
go test ./... -bench=.
```