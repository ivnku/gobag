package maps

// GetKeys returns a slice of keys of a map
func GetKeys[K comparable, V any](in map[K]V) []K {
	if in == nil {
		return nil
	}

	out := make([]K, 0, len(in))

	for k := range in {
		out = append(out, k)
	}

	return out
}
