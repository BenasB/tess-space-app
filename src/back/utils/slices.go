package utils

func MapFiltered[T, V any](input []T, transform func(int, T) (V, bool)) []V {
	vs := make([]V, 0, len(input))
	for i, t := range input {
		v, ok := transform(i, t)
		if ok {
			vs = append(vs, v)
		}
	}
	return vs
}
