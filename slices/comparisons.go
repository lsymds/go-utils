package slices

// Min finds the minimum value in a slice of numbers.
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

	m := slice[0]

	for _, i := range slice {
		if comparer(&i) < comparer(&m) {
			m = i
		}
	}

	return &m, nil
}

// Max finds the maximum value in a slice of numbers.
func Max[T ordered](slice []T) (T, error) {
	if len(slice) == 0 {
		return *new(T), ErrEmptySlice
	}

	m := slice[0]

	for _, i := range slice {
		if i > m {
			m = i
		}
	}

	return m, nil
}

// MaxBy finds the maximum value in a slice by comparing the value returned from the comparer
// closure for the given slice element.
func MaxBy[T any, TComparer ordered](slice []T, comparer func(el *T) TComparer) (*T, error) {
	if len(slice) == 0 {
		return new(T), ErrEmptySlice
	}

	m := slice[0]

	for _, i := range slice {
		if comparer(&i) > comparer(&m) {
			m = i
		}
	}

	return &m, nil
}
