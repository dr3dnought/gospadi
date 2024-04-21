package gospadi

func Reverse[E any](arr []E) []E {
	l := len(arr)
	if l == 1 {
		return append(arr)
	}

	result := make([]E, 0, l)
	for i := len(arr) - 1; i >= 0; i-- {
		result = append(result, arr[i])
	}

	return result
}
