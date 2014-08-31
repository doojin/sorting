Go Sorting
==========

Array sorting algorithms implemented in Go

Algorithms implemented:
- Bubble sorting
- Selection sorting
- Coctail sorting
- Gnome sorting
- Insertion sorting
- Quick sorting
- Merge sorting
- Tree sorting

Examples:

	package main

	import (
		"./sorting"
		"fmt"
	)

	func main() {
		arr := []int{13, -45, 28, 0, 112, -5}

		result := sorting.BubbleSort(arr)
		fmt.Println(result)

		result = sorting.InsertionSort(arr)
		fmt.Println(result)

		result = sorting.GnomeSort(arr)
		fmt.Println(result)

		result = sorting.TreeSort(arr)
		fmt.Println(result)

		// and others...
	}
