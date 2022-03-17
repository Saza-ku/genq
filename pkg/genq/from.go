package genq

func From[T any](src *[]T) Query[T] {
	return Query[T]{
		Iterate: func() Iterator[T] {
			len := len(*src)
			index := 0

			return func() (item T, ok bool) {
				ok = index < len
				if ok {
					item = (*src)[index]
					index++
				}
				return
			}
		},
	}
}
