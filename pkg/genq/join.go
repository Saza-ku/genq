package genq

type Pair[X, Y any] struct {
	First  X
	Second Y
}

func Join[X, Y any](joinPred func(x X, y Y) bool, qx Query[X], qy Query[Y]) Query[Pair[X, Y]] {
	return Query[Pair[X, Y]]{
		Iterate: func() Iterator[Pair[X, Y]] {
			nextX := qx.Iterate()
			nextY := qy.Iterate()
			itemX, okX := nextX()

			if !okX {
				return func() (item Pair[X, Y], ok bool) {
					return
				}
			}

			return func() (item Pair[X, Y], ok bool) {
				for true {
					itemY, okY := nextY()

					for !okY {
						itemX, okX = nextX()
						if !okX {
							return
						}
						nextY = qy.Iterate()
						itemY, okY = nextY()
					}

					if joinPred(itemX, itemY) {
						item = Pair[X, Y]{First: itemX, Second: itemY}
						ok = true
						return
					}
				}
				return
			}
		},
	}
}
