/*
This file is derived from go-linq and modified.
go-linq : https://github.com/ahmetb/go-linq
*/

package genq

type Summable interface {
	~complex64 | ~complex128 | ~float32 | ~float64 |
		~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

func Sum[T Summable](q Query[T]) T {
	next := q.Iterate()
	var sum T
	for item, ok := next(); ok; item, ok = next() {
		sum += item
	}
	return sum
}
