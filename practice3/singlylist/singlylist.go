package singlylist

type node[T any] struct {
	value T
	next  *node[T]
}

type List[T any] struct {
	first *node[T]
	last  *node[T]
	size  int
}

func (l *List[T]) Add(value T) {
	newNode := &node[T]{value: value}

	if l.first == nil {
		l.first = newNode
		l.last = newNode
	} else {
		l.last.next = newNode
		l.last = newNode
	}

	l.size++
}

func (l *List[T]) Get(index int) (T, bool) {
	var zero T

	if index < 0 || index >= l.size {
		return zero, false
	}

	current := l.first
	for i := 0; i < index; i++ {
		current = current.next
	}

	return current.value, true
}

func (l *List[T]) Remove(index int) bool {
	if index < 0 || index >= l.size {
		return false
	}

	if index == 0 {
		l.first = l.first.next
		if l.size == 1 {
			l.last = nil
		}
		l.size--
		return true
	}

	current := l.first
	for i := 0; i < index-1; i++ {
		current = current.next
	}

	current.next = current.next.next

	if index == l.size-1 {
		l.last = current
	}

	l.size--
	return true
}

func (l *List[T]) Values() []T {
	result := make([]T, 0, l.size)
	current := l.first

	for current != nil {
		result = append(result, current.value)
		current = current.next
	}

	return result
}
