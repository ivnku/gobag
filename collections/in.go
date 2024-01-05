package collections

// In returns true if elem exist in the provided collection otherwise returns false
func In[A comparable](in []A, elem A) bool {
	for _, item := range in {
		if item == elem {
			return true
		}
	}

	return false
}
