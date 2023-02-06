package gensort

func BubbleSort[V Number](array []V) {
	for i := 0; i < len(array)-1; i++ {
		if array[i] > array[i+1] {
			array[i], array[i+1] = array[i+1], array[i]
		}
	}

	if len(array) > 2 {
		BubbleSort(array[:len(array)-1])
	}
}
