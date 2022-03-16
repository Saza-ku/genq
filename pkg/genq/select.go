package genq

func Select[X, Y any](selector func(X) Y, q Query[X]) Query[Y] {
	return Query[Y]{
		Iterate: func() Iterator[Y] {
			next := q.Iterate()

			return func() (item Y, ok bool) {
				it, ok := next()
				if ok {
					item = selector(it)
				}
				return
			}
		},
	}
}
