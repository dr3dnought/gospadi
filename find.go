package gospadi

type FindFunc[E any] func(e E) bool

func Find[E any](arr []E, f FindFunc[E]) (E, bool) {
	for _, e := range arr {
		if f(e) {
			return e, true
		}
	}

	return *new(E), false
}
