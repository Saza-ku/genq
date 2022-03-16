package genq

func Select[X, Y any](q Query[X], selector func(X) Y) Query[Y] {
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
