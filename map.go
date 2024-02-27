package gospadi

type MapFunc[E, T any] func(e E) T

func Map[E, T any](s []E, f MapFunc[E, T]) []T {
	mapped := make([]T, len(s))
	for i, el := range s {
		mapped[i] = f(el)
	}

	return mapped
}
