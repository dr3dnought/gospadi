package gospadi

type SomeFunc[E any] func(e E) bool

func Some[E any](arr []E, f SomeFunc[E]) bool {
	for _, e := range arr {
		if f(e) {
			return true
		}
	}

	return false
}
