package hotels

func Filter[T any](s []T, fn func(T) bool) []T {
	r := make([]T, 0, len(s))
	for _, e := range s {
		if fn(e) {
			r = append(r, e)
		}
	}

	return r
}

func Transform[T, K any](s []T, fn func(T) K) []K {
	if s == nil {
		return nil
	}

	out := make([]K, len(s))
	for i, v := range s {
		out[i] = fn(v)
	}
	return out
}
