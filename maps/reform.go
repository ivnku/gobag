package maps

// Reform reformat map's keys and values from one type to another
func Reform[K1 comparable, K2 comparable, V1 any, V2 any](
	in map[K1]V1,
	keyFn func(K1) K2,
	valFn func(V1) V2,
) map[K2]V2 {
	if in == nil {
		return map[K2]V2{}
	}

	out := make(map[K2]V2, len(in))

	for k, v := range in {
		out[keyFn(k)] = valFn(v)
	}

	return out
}
