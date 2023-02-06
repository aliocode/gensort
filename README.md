### Features

- Sorting of numeric type slices;
- Quick, bubble and insertion sorting;
- Native generics support;
- Covered by fuzzing and benchmarks;

# Gensort

![](https://img.icons8.com/ios/452/generic-sorting.png)


**Usage**
#### Code example
```golang
package main

import (
	"fmt"
	"github.com/vrumg/gensort"
)

func main() {
	slice := []int{5, 9, 0, 4, 2, 8, 3, 7, 6, 1}
	gensort.QuickSort(slice)
	fmt.Println(slice)
}

```

#### Output

```bash
go run main.go

[0 1 2 3 4 5 6 7 8 9]
```


#### Benchmarks
Comparison of bubble, insertion, quick and native go sorting algorithms on 10.000 elements slices.\
OS: Linux\
Arch: amd64\
CPU: AMD Ryzen 7 3700X 8-Core Processor\
```bash
go test -bench=. -benchmem

goos: linux
goarch: amd64
pkg: github.com/vrumg/gensort
cpu: AMD Ryzen 7 3700X 8-Core Processor
BenchmarkBubbleSort-16                16          68920794 ns/op          161920 B/op      10001 allocs/op
BenchmarkInsertionSort-16             27          42216993 ns/op          161920 B/op      10001 allocs/op
BenchmarkQuickSort-16               1508            778599 ns/op          161921 B/op      10001 allocs/op
BenchmarkNativeSort-16              1005           1188826 ns/op          161977 B/op      10003 allocs/op
```
