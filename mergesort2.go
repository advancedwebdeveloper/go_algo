package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	slice := generateSlice(20)
	fmt.Println("\n--- Unsorted --- \n\n", slice)
	fmt.Println("\n--- Sorted ---\n\n", mergeSort(slice), "\n")
}

// Generates a slice of size, size filled with random numbers
func generateSlice(size int) []int {

	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}

func mergeSort(items []int) []int {
	var num = len(items)

	if num == 1 {
		return items
	}

	middle := int(num / 2)
	var (
		left  = make([]int, middle)
		right = make([]int, num-middle)
	)
	for i := 0; i < num; i++ {
		if i < middle {
			left[i] = items[i]
		} else {
			right[i-middle] = items[i]
		}
	}

	return merge(mergeSort(left), mergeSort(right))
}

func merge(left, right []int) (result []int) {
	result = make([]int, len(left)+len(right))

	i := 0
	for len(left) > 0 && len(right) > 0 {
		fmt.Println("Until the lengthes of both left and right array are grater then zero")
		if left[0] < right[0] {
			result[i] = left[0]
			left = left[1:]
			fmt.Println(left)
		} else {
			result[i] = right[0]
			right = right[1:]
			fmt.Println(right)
		}
		i++
	}

	for j := 0; j < len(left); j++ {
		fmt.Print("\ni,j values:")
		fmt.Print(i)
		fmt.Print(" ")
		fmt.Print(j)
		result[i] = left[j]
		i++
	}
	for j := 0; j < len(right); j++ {
		fmt.Print("\nresulting array, index at ")
		fmt.Print(i)
		result[i] = right[j]
		fmt.Print("iterator i has it's value changed from ")
		fmt.Print(i)
		i++
		fmt.Print(" to ")
		fmt.Print(i)
	}

	return
}
