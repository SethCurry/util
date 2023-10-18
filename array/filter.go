package array

func RemoveZeros[T comparable](item T) bool {
	var zeroVal T

	return item != zeroVal
}

func Filter[T any](fn func(T) bool, items []T) []T {
	ret := []T{}

	for _, v := range items {
		if fn(v) {
			ret = append(ret, v)
		}
	}

	return ret
}

func First[T any](fn func(T) bool, items []T) (T, bool) {
	for _, v := range items {
		if fn(v) {
			return v, true
		}
	}

	var zeroVal T

	return zeroVal, false
}

func Contains[T comparable](items []T, item T) bool {
	for _, v := range items {
		if v == item {
			return true
		}
	}

	return false
}
