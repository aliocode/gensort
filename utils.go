package gensort

import (
	"github.com/pkg/errors"
	"math/rand"
	"reflect"
	"time"
)

func getParametrizedSliceToSort[V Number](length int, min V, max V) ([]V, error) {
	toSort := make([]V, 0, length)

	rand.Seed(time.Now().UnixNano())
	for len(toSort) < length {
		number, err := getNumber[V](min, max-1)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to generate number by type %s", reflect.TypeOf(min).String())
		}
		toSort = append(toSort, number)
	}

	return toSort, nil
}

func getNumber[V Number](min V, max V) (V, error) {
	switch reflect.TypeOf(min).String() {
	case "int", "int8", "int16", "int32", "int64":
		num := V(rand.Intn(int(max)-int(min)+1) + int(min))
		return num, nil
	case "uint", "uint8", "uint16", "uint32", "uint64", "byte":
		num := V(rand.Uint32()%uint32(max-min)) + min
		return num, nil
	case "float32", "float64":
		num := V(rand.Float32()*float32(max-min)) + min
		return num, nil
	default:
		return 0, errors.New("invalid type")
	}
}

func isSliceSorted[V Number](slice []V) bool {
	if len(slice) == 0 {
		return false
	}

	if len(slice) == 1 {
		return true
	}

	previous := slice[0]
	for i := 1; i < len(slice); i++ {
		if previous > slice[i] {
			return false
		}
		previous = slice[i]
	}

	return true
}
