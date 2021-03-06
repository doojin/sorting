package sorting

import "math"

func StupidSort(xs []int) []int {
	i := 0
	for i < len(xs)-1 {
		if xs[i] > xs[i+1] {
			xs[i], xs[i+1] = xs[i+1], xs[i]
			i = 0
		} else {
			i++
		}
	}
	return xs
}

// ****************************************

func BubbleSort(xs []int) []int {
	for i := 0; i < len(xs)-1; i++ {
		for j := 0; j < len(xs)-i-1; j++ {
			if xs[j] > xs[j+1] {
				xs[j], xs[j+1] = xs[j+1], xs[j]
			}
		}
	}
	return xs
}

// ****************************************

func SelectionSort(xs []int) []int {
	for i := 0; i < len(xs)-1; i++ {
		indexMin := i
		for j := i + 1; j < len(xs); j++ {
			if xs[j] < xs[indexMin] {
				indexMin = j
			}
		}
		xs[i], xs[indexMin] = xs[indexMin], xs[i]
	}
	return xs
}

// ****************************************

func CoctailSort(xs []int) []int {
	swapped := true

	for swapped {
		swapped = false

		for i := 0; i < len(xs)-2; i++ {
			if xs[i] > xs[i+1] {
				xs[i], xs[i+1] = xs[i+1], xs[i]
				swapped = true
			}
		}

		if !swapped {
			return xs
		}
		swapped = false

		for i := len(xs) - 2; i >= 0; i-- {
			if xs[i] > xs[i+1] {
				xs[i], xs[i+1] = xs[i+1], xs[i]
				swapped = true
			}
		}
	}
	return xs
}

// ****************************************

func GnomeSort(xs []int) []int {
	i := 1
	for i < len(xs) {
		if i == 0 {
			i = 1
		}

		if xs[i-1] <= xs[i] {
			i++
		} else {
			xs[i], xs[i-1] = xs[i-1], xs[i]
			i--
		}
	}
	return xs
}

// ****************************************

func InsertionSort(xs []int) []int {
	for i := 1; i < len(xs); i++ {
		prevIndex := i - 1
		current := xs[i]
		for prevIndex >= 0 && xs[prevIndex] > current {
			xs[prevIndex+1] = xs[prevIndex]
			xs[prevIndex] = current
			prevIndex--
		}
	}
	return xs
}

// ****************************************

func QuickSort(xs []int, start int, end int) []int {
	partition := func(xs []int, start int, end int) int {
		mark := start
		for i := start; i <= end; i++ {
			if xs[i] <= xs[end] {
				xs[mark], xs[i] = xs[i], xs[mark]
				mark++
			}
		}
		return mark - 1
	}

	if len(xs) == 0 || start >= end {
		return xs
	}

	pivot := partition(xs, start, end)
	QuickSort(xs, start, pivot-1)
	QuickSort(xs, pivot+1, end)
	return xs
}

// ****************************************

func MergeSort(xs []int) []int {
	// Merging 2 slices into the one (sorted)
	merge := func(xs1 []int, xs2 []int) []int {
		sorted := make([]int, len(xs1)+len(xs2))
		var i1, i2, curr int
		for i1 < len(xs1) || i2 < len(xs2) {
			if i1 < len(xs1) && i2 < len(xs2) {
				if xs1[i1] < xs2[i2] {
					sorted[curr] = xs1[i1]
					i1++
				} else {
					sorted[curr] = xs2[i2]
					i2++
				}
				curr++
			} else {
				if i1 < len(xs1) {
					sorted[curr] = xs1[i1]
					i1++
				}
				if i2 < len(xs2) {
					sorted[curr] = xs2[i2]
					i2++
				}
				curr++
			}
		}
		return sorted
	}

	if len(xs) < 2 {
		return xs
	}
	middle := int(math.Ceil(float64(len(xs) / 2)))
	leftSlice := xs[0:middle]
	rightSlice := xs[middle:]
	return merge(MergeSort(leftSlice), MergeSort(rightSlice))
}

// ****************************************

func TreeSort(xs []int) []int {

	// Structure, which the binary tree will be made of
	type Tree struct {
		left  *Tree
		right *Tree
		value int
	}
	var insert func(node *Tree, tree *Tree)
	var traverse func(tree *Tree)

	root := new(Tree)
	result := make([]int, 0)

	// Inserting new element to the left/right depending on it's value
	insert = func(node *Tree, tree *Tree) {
		if node.value > tree.value {
			if node.left == nil {
				node.left = tree
			} else {
				insert(node.left, tree)
			}
		} else {
			if node.right == nil {
				node.right = tree
			} else {
				insert(node.right, tree)
			}
		}
	}

	// Transform binary tree to result slice
	traverse = func(tree *Tree) {
		if tree.left != nil {
			traverse(tree.left)
		}
		result = append(result, tree.value)
		if tree.right != nil {
			traverse(tree.right)
		}
	}

	if len(xs) == 0 {
		return xs
	}

	root.value = xs[0]
	for i := 1; i < len(xs); i++ {
		insert(root, &Tree{value: xs[i]})
	}
	traverse(root)
	return result
}
