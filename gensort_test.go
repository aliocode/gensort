package gensort

import (
	"sort"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	toSort, err := getParametrizedSliceToSort[int](100, -10, 10)
	if err != nil {
		t.Errorf("%v", err)
	}

	BubbleSort[int](toSort)
	if !isSliceSorted(toSort) {
		t.Error()
	}
}

func TestInsertionSort(t *testing.T) {
	toSort, err := getParametrizedSliceToSort[int](100, -10, 10)
	if err != nil {
		t.Errorf("%v", err)
	}

	InsertionSort[int](toSort)
	if !isSliceSorted(toSort) {
		t.Error()
	}
}

func TestQuickSort(t *testing.T) {
	toSort, err := getParametrizedSliceToSort[int](100, -10, 10)
	if err != nil {
		t.Errorf("%v", err)
	}

	QuickSort(toSort)
	if !isSliceSorted(toSort) {
		t.Error()
	}
}

func TestNativeSort(t *testing.T) {
	toSort, err := getParametrizedSliceToSort[int](100, -10, 10)
	if err != nil {
		t.Errorf("%v", err)
	}

	sort.Slice(toSort, func(i, j int) bool {
		return toSort[i] < toSort[j]
	})
	if !isSliceSorted(toSort) {
		t.Error()
	}
}

func TestGetParametrizedSliceToSort(t *testing.T) {
	t.Run("type int", func(t *testing.T) {
		length := 20
		min := -10
		max := 10
		toSort, err := getParametrizedSliceToSort[int](length, min, max)
		if err != nil {
			t.Errorf("%v", err)
		}
		if toSort == nil {
			t.Error()
		}
		t.Log(toSort)
	})

	t.Run("type byte", func(t *testing.T) {
		length := 20
		min := byte(0)
		max := byte(10)
		toSort, err := getParametrizedSliceToSort[byte](length, min, max)
		if err != nil {
			t.Errorf("%v", err)
		}
		if toSort == nil {
			t.Error()
		}
		t.Log(toSort)
	})

	t.Run("type uint8", func(t *testing.T) {
		length := 20
		min := uint8(0)
		max := uint8(10)
		toSort, err := getParametrizedSliceToSort[uint8](length, min, max)
		if err != nil {
			t.Errorf("%v", err)
		}
		if toSort == nil {
			t.Error()
		}
		t.Log(toSort)
	})

	t.Run("type float32", func(t *testing.T) {
		length := 200
		min := float32(-1.0)
		max := float32(1.0)
		toSort, err := getParametrizedSliceToSort[float32](length, min, max)
		if err != nil {
			t.Errorf("%v", err)
		}
		if toSort == nil {
			t.Error()
		}
		t.Log(toSort)
	})
}

func FuzzBubbleSort(f *testing.F) {
	f.Add([]byte{5, 9, 0, 4, 2, 8, 3, 7, 6, 1})
	f.Fuzz(func(t *testing.T, slice []byte) {
		BubbleSort[byte](slice)
	})
}
func FuzzInsertionSort(f *testing.F) {
	f.Add([]byte{5, 9, 0, 4, 2, 8, 3, 7, 6, 1})
	f.Fuzz(func(t *testing.T, slice []byte) {
		InsertionSort[byte](slice)
	})
}

func FuzzQuickSort(f *testing.F) {
	f.Add([]byte{5, 9, 0, 4, 2, 8, 3, 7, 6, 1})
	f.Fuzz(func(t *testing.T, toSort []byte) {
		QuickSort[byte](toSort)
	})
}

func BenchmarkBubbleSort(b *testing.B) {
	for n := 0; n < b.N; n++ {
		toSort, err := getParametrizedSliceToSort[int](10_000, -10_000, 10_000)
		if err != nil {
			b.Errorf("%v", err)
		}

		BubbleSort[int](toSort)
	}
}

func BenchmarkInsertionSort(b *testing.B) {
	for n := 0; n < b.N; n++ {
		toSort, err := getParametrizedSliceToSort[int](10_000, -10_000, 10_000)
		if err != nil {
			b.Errorf("%v", err)
		}

		InsertionSort[int](toSort)
	}
}

func BenchmarkQuickSort(b *testing.B) {
	for n := 0; n < b.N; n++ {
		toSort, err := getParametrizedSliceToSort[int](10_000, -10_000, 10_000)
		if err != nil {
			b.Errorf("%v", err)
		}

		QuickSort(toSort)
	}
}

func BenchmarkNativeSort(b *testing.B) {
	for n := 0; n < b.N; n++ {
		toSort, err := getParametrizedSliceToSort[int](10_000, -10_000, 10_000)
		if err != nil {
			b.Errorf("%v", err)
		}

		sort.Slice(toSort, func(i, j int) bool {
			return toSort[i] < toSort[j]
		})
	}
}
