package slices

// Map projects each element of a slice and returns a new array consisting of the projected values.
func Map[T any, TResult any](slice []T, projection func(*T) TResult) ([]TResult, error) {
	if len(slice) == 0 {
		return nil, ErrEmptySlice
	}

	result := make([]TResult, 0)

	for _, el := range slice {
		result = append(result, projection(&el))
	}

	return result, nil
}
