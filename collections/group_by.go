package collections

// GroupBy grouping a collection of type I by the key of type K
// which is formed by the keyFn(I) K function and with the value of type []V
// each item of which is formed by the valueFn(I) V function
func GroupBy[I any, K comparable, V any](
	in []I,
	keyFn func(I) K,
	valueFn func(I) V,
) map[K][]V {
	out := make(map[K][]V, len(in))

	for _, v := range in {
		k := keyFn(v)
		out[k] = append(out[k], valueFn(v))
	}

	return out
}
