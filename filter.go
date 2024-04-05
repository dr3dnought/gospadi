package gospadi

type FilterFunc[E any] func(e E) bool

func Filter[E any](s []E, f FilterFunc[E]) (res []E) {
	for _, e := range s {
		if f(e) {
			res = append(res, e)
		}
	}

	return
}
