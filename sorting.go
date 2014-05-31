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