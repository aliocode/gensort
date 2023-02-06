package gensort

func QuickSort[V Number](slice []V) {
	if len(slice) < 1 {
		return
	}
	p := partition[V](slice)

	if p-1 > 0 {
		QuickSort[V](slice[:p])
	}
	if p+1 < len(slice)-1 {
		QuickSort[V](slice[p+1:])
	}

}

func partition[V Number](slice []V) int {
	high := len(slice) - 1
	pivot := slice[high]
	i := 0
	for j := 0; j < high; j++ {
		if slice[j] < pivot {
			slice[i], slice[j] = slice[j], slice[i]
			i++
		}
	}

	slice[i], slice[high] = slice[high], slice[i]

	return i
}
