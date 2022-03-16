package genq

func From[T any](src []T) Query[T] {
	len := len(src)
	return Query[T]{
		Iterate: func() Iterator[T] {
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
