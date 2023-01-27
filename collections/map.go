package collections

// Map transforms collection of type A into collection
// of type B with mapping function fn(A) B
func Map[A any, B any](in []A, fn func(A) B) []B {
	out := make([]B, 0, len(in))

	for _, item := range in {
		out = append(out, fn(item))
	}

	return out
}
