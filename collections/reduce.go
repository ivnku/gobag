package collections

// Reduce accumulates the data from collection of type []A with
// the function fn func(acc B, item A) B
func Reduce[A any, B any](in []A, acc B, fn func(acc B, item A) B) B {
	for _, item := range in {
		acc = fn(acc, item)
	}

	return acc
}
