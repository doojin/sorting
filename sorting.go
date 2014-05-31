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

func InsertionSort(xs []int) []int {
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