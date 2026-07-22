package hotels

func GroupBy[T any, K comparable](s []T, fn func(e T) K) map[K][]T {
	if s == nil {
		return nil
	}

	result := map[K][]T{}
	for _, e := range s {
		key := fn(e)
		result[key] = append(result[key], e)
	}

	return result
}

func Filter[T any](s []T, fn func(T) bool) []T {
	if s == nil {
		return nil
	}

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
