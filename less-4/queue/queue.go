package queue

type Queue[T any] struct {
	s    []T
	low  int
	high int
	size int
}

func New[T any](size int) *Queue[T] {
	return &Queue[T]{
		s:    make([]T, size),
		size: size,
		low:  -1,
		high: -1,
	}
}

func (q *Queue[T]) Push(v T) {
	if q.high >= q.size-1 {
		return
	}
	q.high++
	q.s[q.high] = v
	if q.low == -1 {
		q.low = 0
	}
}

func (q *Queue[T]) Pop() T {
	var zero T
	if q.low == -1 || q.low > q.high {
		return zero
	}
	val := q.s[q.low]
	q.low++
	return val
}

func (q *Queue[T]) IsEmpty() bool {
	return q.low == -1 || q.low > q.high
}

func (q *Queue[T]) Size() int {
	if q.IsEmpty() {
		return 0
	}
	return q.high - q.low + 1
}
