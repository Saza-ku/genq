/*
This file is derived from go-linq and modified.
go-linq : https://github.com/ahmetb/go-linq
*/

package genq

type Iterator[T any] func() (item T, ok bool)

type Query[T any] struct {
	Iterate func() Iterator[T]
}

func (q Query[T]) ToSlice() []T {
	next := q.Iterate()
	slice := make([]T, 0)
	for item, ok := next(); ok; item, ok = next() {
		slice = append(slice, item)
	}
	return slice
}

func (q Query[T]) First() (first T, ok bool) {
	next := q.Iterate()
	return next()
}

func (q Query[T]) Any() bool {
	_, ok := q.Iterate()()
	return ok
}

func (q Query[T]) AnyWith(predicate func(T) bool) bool {
	next := q.Iterate()

	for item, ok := next(); ok; item, ok = next() {
		if predicate(item) {
			return true
		}
	}

	return false
}
