/*
This file is derived from go-linq and modified.
go-linq : https://github.com/ahmetb/go-linq
*/

package genq

func Where[T any](predicate func(T) bool, q Query[T]) Query[T] {
	return Query[T]{
		Iterate: func() Iterator[T] {
			next := q.Iterate()

			return func() (item T, ok bool) {
				for item, ok = next(); ok; item, ok = next() {
					if predicate(item) {
						return
					}
				}

				return
			}
		},
	}
}
