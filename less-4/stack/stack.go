package stack

type Stack[T any] struct {
	a    []T
	head int
}

func New[T any](size int) *Stack[T] {
	return &Stack[T]{
		a:    make([]T, size),
		head: -1,
	}
}

func (s *Stack[T]) Push(v T) {
	s.head++
	s.a[s.head] = v
}

func (s *Stack[T]) Pop() T {
	var zero T
	if s.head < 0 {
		return zero
	}
	v := s.a[s.head]
	s.head--
	return v
}

func (s *Stack[T]) Peek() T {
	var zero T
	if s.head < 0 {
		return zero
	}
	return s.a[s.head]
}

func (s *Stack[T]) IsEmpty() bool {
	return s.head < 0
}

func (s *Stack[T]) Size() int {
	return s.head + 1
}
