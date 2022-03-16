package genq

type Group[K, V any] struct {
	Key   K
	Group []V
}

func GroupBy[K comparable, V any](keySelector func(V) K, q Query[V]) Query[Group[K, V]] {
	return Query[Group[K, V]]{
		func() Iterator[Group[K, V]] {
			next := q.Iterate()
			set := make(map[K][]V)

			for item, ok := next(); ok; item, ok = next() {
				key := keySelector(item)
				set[key] = append(set[key], item)
			}

			len := len(set)
			idx := 0
			groups := make([]Group[K, V], len)
			for k, v := range set {
				groups[idx] = Group[K, V]{k, v}
				idx++
			}

			index := 0

			return func() (item Group[K, V], ok bool) {
				ok = index < len
				if ok {
					item = groups[index]
					index++
				}

				return
			}
		},
	}
}
