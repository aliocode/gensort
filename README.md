# Gensort
An experiment with sorting algorithms and generics.

## Features
- Sorting of numeric type slices.
- Quick, bubble and insertion sorting.
- Native generics support.
- Covered by fuzzing and benchmarks.

## Usage example

### Source code
```golang
package main

import (
	"fmt"
	"github.com/aliocode/gensort"
)

func main() {
	slice := []int{5, 9, 0, 4, 2, 8, 3, 7, 6, 1}
	gensort.QuickSort(slice)
	fmt.Println(slice)
}

```

### Output
```bash
go run main.go

[0 1 2 3 4 5 6 7 8 9]
```

## Benchmarks
Comparison of a bubble, insertion, quick and native go sorting algorithms on 10.000 elements slices.

```bash
BenchmarkBubbleSort-16                16          68920794 ns/op          161920 B/op      10001 allocs/op
BenchmarkInsertionSort-16             27          42216993 ns/op          161920 B/op      10001 allocs/op
BenchmarkQuickSort-16               1508            778599 ns/op          161921 B/op      10001 allocs/op
BenchmarkNativeSort-16              1005           1188826 ns/op          161977 B/op      10003 allocs/op
```
