package collections

// Filter filters collection of type A by condition cond()
func Filter[A any](in []A, cond func(A) bool) []A {
	out := make([]A, 0, len(in))

	for _, item := range in {
		if cond(item) {
			out = append(out, item)
		}
	}

	return out
}