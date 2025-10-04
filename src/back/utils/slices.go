package utils

func MapFiltered[T, V any](input []T, transform func(T) (V, bool)) []V {
	vs := make([]V, 0, len(input))
	for _, t := range input {
		v, ok := transform(t)
		if ok {
			vs = append(vs, v)
		}
	}
	return vs
}
