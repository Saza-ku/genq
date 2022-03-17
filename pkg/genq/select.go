/*
This file is derived from go-linq and modified.
go-linq : https://github.com/ahmetb/go-linq
*/

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

func SelectMany[X, Y any](selector func(X) []Y, q Query[X]) Query[Y] {
	return Query[Y]{
		Iterate: func() Iterator[Y] {
			next := q.Iterate()
			var items []Y
			index := 0

			return func() (item Y, ok bool) {
				for index >= len(items) {
					var it X
					it, ok = next()
					if !ok {
						return
					}
					items = selector(it)
					index = 0
				}
				item = items[index]
				ok = true
				index++

				return
			}
		},
	}
}
