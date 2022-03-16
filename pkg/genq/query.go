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
