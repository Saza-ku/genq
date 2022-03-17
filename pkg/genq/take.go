package genq

func Take[T any](count int, q Query[T]) Query[T] {
	return Query[T]{
		Iterate: func() Iterator[T] {
			next := q.Iterate()
			n := count

			return func() (item T, ok bool) {
				if n <= 0 {
					return
				}

				n--
				return next()
			}
		},
	}
}
