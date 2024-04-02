package apply

func Slice[T any, S any](in []T, f func(T) S) (out []S) {
	for _, elem := range in {
		out = append(out, f(elem))
	}
	return
}
