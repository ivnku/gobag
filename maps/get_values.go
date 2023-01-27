package maps

// GetValues returns a slice of values of a map
func GetValues[K comparable, V any](in map[K]V) []V {
	if in == nil {
		return nil
	}

	out := make([]V, 0, len(in))

	for _, v := range in {
		out = append(out, v)
	}

	return out
}
