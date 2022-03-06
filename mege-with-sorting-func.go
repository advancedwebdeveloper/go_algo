package main

import "fmt"

func main() {
	fmt.Println("\nascending:", merge([]int{1, 3, 6}, []int{2, 4, 5}, func(a, b int) bool {
		// ascending
		return a < b
	}))
	fmt.Println("\ndescending:", merge([]int{1, 3, 6}, []int{2, 4, 5}, func(a, b int) bool {
		// descending
		return a > b
	}))
}

type sortFunc func(a, b int) bool

// merge merges 2 slices into a single slice, also sorting the resulting slice
// Note: to ensure the merge function works correctly the L and R must be sorted
func merge(L, R []int, fn sortFunc) []int {
	A := make([]int, len(L)+len(R))
	i, j, k := 0, 0, 0
	// compare left & right side elements before merging
	for i < len(L) && j < len(R) {
		if fn(L[i], R[j]) {
			fmt.Println("\nfn(L[i], R[j]) returned true")
			A[k] = L[i]
			fmt.Print("A[")
			fmt.Print(k)
			fmt.Print("] = R[")
			fmt.Print(j)
			fmt.Print("]")
			fmt.Println("\ni changed from ")
			fmt.Print(i)
			fmt.Print(" to ")
			i++
			fmt.Print(i)
		} else {
			fmt.Print("\nA[")
			fmt.Print(k)
			fmt.Print("] = R[")
			fmt.Print(j)
			fmt.Print("]")
			A[k] = R[j]
			j++
		}
		k++
	}

	// check if any elements from the left/right side
	// were missed in the comparison section above
	for i < len(L) {
		A[k] = L[i]
		i++
		k++
	}
	for j < len(R) {
		A[k] = R[j]
		j++
		k++
	}

	return A
}
