package main

import "fmt"

// This is the only exported function which will take a slice as an input
// and return a different slice with the original integers sorted
func Sort(input []int) []int {
	// making a new slice and copying the contents of the input
	// into it
	// this slice is going to be subsequently mutated so that
	// we get the desired order of elements
	sorted := make([]int, len(input))
	copy(sorted, input)

	// calling the private "subSort" function which can sort a
	// sub-slice, by taking two more arguments: the start and
	// end index
	subSort(sorted, 0, len(input)-1)
	return sorted
}

// this function is going to call itself recursively with the left
// and the right halves of the input slice (the divide)
// then call the "merge" function (the conquest)
func subSort(sorted []int, leftStart int, rightEnd int) {
	// stop the recursion if there is nothing to divide
	if leftStart >= rightEnd {
		return
	}
	// calculating the middle element so that we can divide our
	// slice
	middle := (leftStart + rightEnd) / 2
	// calling itself recursively with both halves
	subSort(sorted, leftStart, middle)
	subSort(sorted, middle+1, rightEnd)
	// merging the two sorted halves
	merge(sorted, leftStart, rightEnd)
}

func merge(sorted []int, leftStart int, rightEnd int) {
	fmt.Println("creating a temporary slice, as we can't easily do it in place")
	fmt.Println("len(sorted) equals ")
	fmt.Print(len(sorted))
	temp := make([]int, len(sorted))
	copy(temp, sorted)

	fmt.Println(" end of the left sub-slice will end in the middle")

	leftEnd := (leftStart + rightEnd) / 2
	fmt.Print("leftEnd equals ")
	fmt.Print(leftEnd)
	fmt.Println(" the start of the right slice will be right after the left ends")
	rightStart := leftEnd + 1
	fmt.Print("rightStart equals ")
	fmt.Print(rightStart)

	left := leftStart
	right := rightStart
	index := leftStart

	fmt.Println(" this is the loop where the actual sorting happens")
	fmt.Println(" we iterate until either the left or the right sub-slice")
	fmt.Println(" runs out of elements")
	for left <= leftEnd && right <= rightEnd {
		fmt.Println(" here we start adding elements to the temporary")
		fmt.Println(" slice we created above")
		fmt.Println(" we choose the smaller element every time")
		if sorted[left] < sorted[right] {
			fmt.Print("sorted[")
			fmt.Print(left)
			fmt.Print("] < sorted[")
			fmt.Print(right)
			fmt.Print("]")
			temp[index] = sorted[left]
			left++
		} else {
			temp[index] = sorted[right]
			right++
		}
		index++
	}

	// here we append to the temporary slice the remaining elements
	// that were not picked in the loop above
	// first we do it for the left and the for the right sub-slice
	// one of them will not contain any remaining elements
	// so will make no changes
	copy(temp[index:rightEnd+1], sorted[right:rightEnd+1])
	copy(temp[index:rightEnd+1], sorted[left:leftEnd+1])
	// finally we store the sorted numbers from the temporary slice
	// into our sorted slice
	copy(sorted, temp)
}
