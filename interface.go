package gensort

type Number interface {
	int | int8 | int16 | int32 | int64 | uint | uint16 | uint32 | uint64 | float32 | float64 | byte
}
