package queue

type Queue[T any] struct {
	items []T
}

func (q *Queue[T]) Push(value T) {
	q.items = append(q.items, value)
}

func (q *Queue[T]) Pop() (T, bool) {
	var zero T

	if len(q.items) == 0 {
		return zero, false
	}

	value := q.items[0]
	q.items = q.items[1:]

	return value, true
}

func (q *Queue[T]) Values() []T {
	return q.items
}
