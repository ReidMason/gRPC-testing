package arrays

func Filter[T any](arr []T, fn func(x T) bool) []T {
	res := make([]T, 0)
	for _, x := range arr {
		if fn(x) {
			res = append(res, x)
		}
	}

	return res
}

func Map[T, Y any](arr []T, fn func(x T) Y) []Y {
	res := make([]Y, 0, len(arr))
	for _, x := range arr {
		res = append(res, fn(x))
	}

	return res
}

func FirstOrDefault[T any](arr []T, defaultValue T) T {
	if len(arr) > 0 {
		return arr[0]
	}

	return defaultValue
}
