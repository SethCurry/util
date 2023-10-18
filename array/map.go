package array

func Map[T any, P any](items []T, fn func(T) P) []P {
	results := make([]P, len(items))

	for k, v := range items {
		results[k] = fn(v)
	}

	return results
}

func MapE[T any, P any](items []T, fn func(T) (P, error)) ([]P, error) {
	results := make([]P, len(items))

	for k, v := range items {
		r, err := fn(v)
		if err != nil {
			return results, err
		}

		results[k] = r
	}

	return results, nil
}
