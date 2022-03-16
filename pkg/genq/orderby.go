package genq

import "sort"

type Ordered interface {
	~string | ~float32 | ~float64 |
		~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

func OrderBy[T any, O Ordered](orderSelector func(T) O, q Query[T]) Query[T] {
	return Query[T]{
		Iterate: func() Iterator[T] {
			src := q.ToSlice()
			sort.Slice(src, func(i, j int) bool { return orderSelector(src[i]) < orderSelector(src[j]) })

			len := len(src)
			index := 0

			return func() (item T, ok bool) {
				ok = index < len
				if ok {
					item = src[index]
					index++
				}
				return
			}
		},
	}
}

func OrderByDescending[T any, O Ordered](orderSelector func(T) O, q Query[T]) Query[T] {
	return Query[T]{
		Iterate: func() Iterator[T] {
			src := q.ToSlice()
			sort.Slice(src, func(i, j int) bool { return orderSelector(src[i]) > orderSelector(src[j]) })

			len := len(src)
			index := 0

			return func() (item T, ok bool) {
				ok = index < len
				if ok {
					item = src[index]
					index++
				}
				return
			}
		},
	}
}
