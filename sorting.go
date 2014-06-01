package sorting

func BubbleSort(xs []int) []int {
	for i:=0; i<len(xs)-1; i++ {
		for j:=0; j<len(xs)-i-1; j++ {
			if xs[j] > xs[j+1] {
				temp := xs[j]
				xs[j] = xs[j+1]
				xs[j+1] = temp
			}
		}
	}
	return xs
}

// ****************************************

func SelectionSort(xs []int) []int {
	for i:=0; i<len(xs)-1; i++ {
		indexMin := i
		for j:=i+1; j<len(xs); j++ {
			if xs[j] < xs[indexMin] {
				indexMin = j
			}
		}
		temp := xs[i]
		xs[i] = xs[indexMin]
		xs[indexMin] = temp
	}
	return xs
}

// ****************************************

func CoctailSort(xs []int) []int {
	swapped := true
	
	for swapped {
		swapped = false
		
		for i:=0; i<len(xs)-2; i++ {
			if xs[i] > xs[i+1] {
				temp := xs[i]
				xs[i] = xs[i+1]
				xs[i+1] = temp
				swapped = true
			}
		}
		
		if !swapped {
			return xs
		}
		swapped = false
		
		for i:=len(xs)-2; i>=0; i-- {
			if xs[i] > xs[i+1] {
				temp := xs[i]
				xs[i] = xs[i+1]
				xs[i+1] = temp
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
			temp := xs[i]
			xs[i] = xs[i-1]
			xs[i-1] = temp
			i--
		}
	}
	return xs
}

// ****************************************

func InsertionSort(xs []int) []int {
	for i:=1; i<len(xs); i++ {
		prevIndex := i-1
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
		for i:=start; i<=end; i++ {
			if xs[i] <= xs[end] {
				temp := xs[mark]
				xs[mark] = xs[i]
				xs[i] = temp
				mark++
			}
		}
		return mark-1
	}
	
	if len(xs) == 0 || start >= end{
		return xs
	}
	
	pivot := partition(xs, start, end)
	QuickSort(xs, start, pivot-1)
	QuickSort(xs, pivot+1, end)
	return xs
}