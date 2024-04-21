package gospadi

func Pop[E any](arr []E) *E {
	l := len(arr)
	if l == 0 {
		return nil
	}

	return &arr[l-1]
}
