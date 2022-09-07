package slices

// integer is a constraint for all types and derived types that are fundamentally an integer.
type integer interface {
	~int8 | ~uint8 | ~int16 | ~uint16 | ~int32 | ~uint32 | ~int | ~uint | ~int64 | ~uint64
}

// float is a constraint for all types and derived types that are fundamentally a float.
type float interface {
	~float32 | ~float64
}

// ordered is a union constraint for all types and derived types that can be compared by order
// operators (<, >, etc).
type ordered interface {
	integer | float
}
