package gospadi

type MapErrFunc[E, T any] func(e E) (T, error)

func MapErr[E, T any](s []E, f MapErrFunc[E, T]) ([]T, error) {
	mapped := make([]T, len(s))
	for i, el := range s {
		n, err := f(el)
		if err != nil {
			return nil, err
		}
    mapped[i] = n
	}
  return mapped, nil
}
