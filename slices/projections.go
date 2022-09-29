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

// FlatMap projects each element of a slice to a slice and then flattens the result, returning a
// new, one dimensional slice of the projected values.
func FlatMap[T any, TResult any](slice []T, projection func(*T) []TResult) ([]TResult, error) {
	if len(slice) == 0 {
		return nil, ErrEmptySlice
	}

	result := make([]TResult, 0)

	for _, el := range slice {
		result = append(result, projection(&el)...)
	}

	return result, nil
}
