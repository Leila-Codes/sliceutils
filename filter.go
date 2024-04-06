package sliceutils

type FilterFunc[T interface{}] func(T) bool

// Filter applies the given FilterFunc against the input slice of T.
// It then returns a new slice of T with only the elements that returned true for the filter.
func Filter[T interface{}](input []T, filter FilterFunc[T]) []T {
	output := make([]T, 0)
	for _, value := range input {
		if filter(value) {
			output = append(output, value)
		}
	}

	return output
}
