package sliceutils

type MapperFunc[IN, OUT interface{}] func(IN) OUT

// Map transforms the given slice of IN type, calling the given MapperFunc for every item.
// The results of all these are returned as a new slice.
func Map[IN, OUT interface{}](input []IN, mapper MapperFunc[IN, OUT]) []OUT {
	output := make([]OUT, len(input))
	for i, value := range input {
		output[i] = mapper(value)
	}

	return output
}
