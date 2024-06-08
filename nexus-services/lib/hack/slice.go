package hack

func Map[InputT any, OutputT any](input []InputT, mapFn func(entry InputT, index int) OutputT) []OutputT {
	output := []OutputT{}
	for index, entry := range input {
		output = append(output, mapFn(entry, index))
	}
	return output
}

type Slice[T any] struct {
	elements []T
}

func NewSlice[T any](initialValues ...T) *Slice[T] {
	return &Slice[T]{elements: []T{}}
}

func (slice *Slice[T]) Append(newElements ...T) *Slice[T] {
	slice.elements = append(slice.elements, newElements...)
	return slice
}

func (slice *Slice[T]) GetElements() []T {
	return slice.elements
}
