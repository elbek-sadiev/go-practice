package linkedlist

type item[T any] struct {
	v    T
	next *item[T]
}

type LinkedList[T any] struct {
	first *item[T]
	last  *item[T]
	size  int
}

func New[T any]() *LinkedList[T] {
	return &LinkedList[T]{}
}

func (l *LinkedList[T]) Add(v T) {
	newItem := &item[T]{v: v}
	if l.size == 0 {
		l.first = newItem
		l.last = newItem
	} else {
		l.last.next = newItem
		l.last = newItem
	}
	l.size++
}

func (l *LinkedList[T]) Get(idx int) T {
	var zero T
	if idx < 0 || idx >= l.size {
		return zero
	}
	current := l.first
	for i := 0; i < idx; i++ {
		current = current.next
	}
	return current.v
}

func (l *LinkedList[T]) Remove(idx int) {
	if idx < 0 || idx >= l.size {
		return
	}
	if idx == 0 {
		l.first = l.first.next
		if l.size == 1 {
			l.last = nil
		}
	} else {
		current := l.first
		for i := 0; i < idx-1; i++ {
			current = current.next
		}
		current.next = current.next.next
		if idx == l.size-1 {
			l.last = current
		}
	}
	l.size--
}

func (l *LinkedList[T]) Values() []T {
	result := make([]T, l.size)
	current := l.first
	for i := 0; i < l.size; i++ {
		result[i] = current.v
		current = current.next
	}
	return result
}

func (l *LinkedList[T]) Size() int {
	return l.size
}
