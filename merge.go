package main

import "fmt"

func main() {
	fmt.Println(merge([]int{1, 3, 6}, []int{2, 4, 5}))
}

func merge(L, R []int) []int {
	fmt.Println("merge merges 2 slices into a single slice, also sorting the resulting slice")
	fmt.Println("Note: to ensure the merge function works correctly the L and R must be sorted")
	A := make([]int, len(L)+len(R))
	fmt.Println("creating a slice A, which has a resulting length, which equals the summ of the L and R slices")
	i, j, k := 0, 0, 0
	// compare left & right side elements before merging
	for i < len(L) && j < len(R) {
		fmt.Println("\nUntil the iterator i is less then the length of of the left slice, and until the iterator j is less then length of the riht slice")
		if L[i] < R[j] {
			fmt.Print("\nAn item L[")
			fmt.Print(i)
			fmt.Print("] is less then R[")
			fmt.Print(j)
			fmt.Print("]")
			A[k] = L[i]
			fmt.Print("\nA[")
			fmt.Print(k)
			fmt.Print("] has an aasigned value \n")
			fmt.Print(L[i])
			i++
		} else { //Uncertain: might be equal or L[i]>R[j]
			A[k] = R[j]
			fmt.Print("A[")
			fmt.Print(k)
			fmt.Print("] = R[")
			fmt.Print(j)
			fmt.Print("]")
			j++
		}
		fmt.Println("\nChanging the value of iterator k from ")
		fmt.Print(k)
		fmt.Print(" to ")
		k++
		fmt.Print(k)
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
