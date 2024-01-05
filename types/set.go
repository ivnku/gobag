package types

type SetInterface[A comparable] interface {
	Add(val A) SetInterface[A]
	IsExists(val A) bool
	Remove(val A) SetInterface[A]
	AddMultiple(values []A) SetInterface[A]
	Len() int
	ToSlice() []A
}

type Set[A comparable] struct {
	storage map[A]struct{}
}

func NewSet[A comparable](items []A) SetInterface[A] {
	set := &Set[A]{storage: make(map[A]struct{}, len(items))}

	if len(items) > 0 {
		return set.AddMultiple(items)
	}

	return set
}

func (s *Set[A]) Add(item A) SetInterface[A] {
	s.storage[item] = struct{}{}

	return s
}

func (s *Set[A]) IsExists(item A) bool {
	_, exist := s.storage[item]

	return exist
}

func (s *Set[A]) Remove(item A) SetInterface[A] {
	delete(s.storage, item)

	return s
}

func (s *Set[A]) AddMultiple(items []A) SetInterface[A] {
	for _, i := range items {
		s.storage[i] = struct{}{}
	}

	return s
}

func (s *Set[A]) Len() int {
	return len(s.storage)
}

func (s *Set[A]) ToSlice() []A {
	slice := make([]A, 0, len(s.storage))

	for item := range s.storage {
		slice = append(slice, item)
	}

	return slice
}
