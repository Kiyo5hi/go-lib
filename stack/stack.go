package stack

type Stack[T any] struct {
	data []T
}

func New[T any](vals ...T) *Stack[T] {
	s := &Stack[T]{
		data: []T{},
	}
	for _, val := range vals {
		s.Push(val)
	}
	return s
}

func (s *Stack[T]) Push(elem T) {
	s.data = append(s.data, elem)
}

func (s *Stack[T]) Pop() (t T, ok bool) {
	if len(s.data) == 0 {
		return t, false
	}

	ind := len(s.data) - 1
	elem := s.data[ind]
	s.data = s.data[:ind]
	return elem, true
}

func (s *Stack[T]) Head() (t T, ok bool) {
	if len(s.data) == 0 {
		return t, false
	}

	return s.data[len(s.data)-1], true
}

func (s *Stack[T]) Len() int {
	return len(s.data)
}
