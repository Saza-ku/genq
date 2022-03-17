/*
This file is derived from go-linq and modified.
go-linq : https://github.com/ahmetb/go-linq
*/

package genq

func Aggregate[X, Y any](f func(X, Y) X, q Query[Y]) (result X) {
	next := q.Iterate()

	for current, ok := next(); ok; current, ok = next() {
		result = f(result, current)
	}

	return result
}
