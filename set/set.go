package set

type Set[T comparable] struct {
	data map[T]struct{}
}

func New[T comparable](vals ...T) *Set[T] {
	s := &Set[T]{
		data: map[T]struct{}{},
	}
	for _, val := range vals {
		s.Add(val)
	}
	return s
}

func (s *Set[T]) Contains(v T) bool {
	_, ok := s.data[v]
	return ok
}

func (s *Set[T]) Add(v T) *Set[T] {
	s.data[v] = struct{}{}
	return s
}

func (s *Set[T]) Remove(v T) *Set[T] {
	delete(s.data, v)
	return s
}

func (s *Set[T]) Union(other *Set[T]) *Set[T] {
	result := New[T]()
	for v := range s.data {
		result.Add(v)
	}
	for v := range other.data {
		result.Add(v)
	}
	return result
}

func (s *Set[T]) Len() int {
	return len(s.data)
}

func (s *Set[T]) ToSlice() []T {
	ts := []T{}
	for t := range s.data {
		ts = append(ts, t)
	}
	return ts
}
