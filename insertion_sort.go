package gensort

func InsertionSort[V Number](array []V) {
	for i := 0; i < len(array)-1; i++ {
		if array[i] > array[i+1] {
			array[i], array[i+1] = array[i+1], array[i]
			// check and swap with sorted head
			for j := i; j >= 1; j-- {
				if array[j] < array[j-1] {
					array[j], array[j-1] = array[j-1], array[j]
					continue
				}
				break
			}

		}
	}
}
