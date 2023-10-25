package slices

import "errors"

// ErrEmptySlice is an error returned when an empty slice is passed to methods.
var ErrEmptySlice = errors.New("cannot perform operations on an empty slice")
