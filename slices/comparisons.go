package slices

// Min finds the minimum value in a slice of ordered numbers.
func Min[T ordered](s []T) (T, error) {
	if len(s) == 0 {
		return *new(T), ErrEmptySlice
	}

	m := s[0]

	for _, i := range s {
		if i < m {
			m = i
		}
	}

	return m, nil
}
