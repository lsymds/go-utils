package slices

// Min finds the minimum value in a slice of ordered numbers.
func Min[T ordered](slice []T) (T, error) {
	if len(slice) == 0 {
		return *new(T), ErrEmptySlice
	}

	m := slice[0]

	for _, i := range slice {
		if i < m {
			m = i
		}
	}

	return m, nil
}

// MinBy finds the minimum value in a slice by comparing the value returned from the comparer
// closure for the given slice element.
func MinBy[T any, TComparer ordered](slice []T, comparer func(el *T) TComparer) (*T, error) {
	if len(slice) == 0 {
		return new(T), ErrEmptySlice
	}

	m := &slice[0]

	for _, i := range slice {
		if comparer(&i) < comparer(m) {
			m = &i
		}
	}

	return m, nil
}
