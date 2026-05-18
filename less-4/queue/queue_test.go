package queue

import "testing"

func TestQueue_PushAndPop(t *testing.T) {
	q := New[int](10)

	q.Push(10)
	q.Push(20)
	q.Push(30)

	if val := q.Pop(); val != 10 {
		t.Errorf("Pop() = %v, want 10", val)
	}
	if val := q.Pop(); val != 20 {
		t.Errorf("Pop() = %v, want 20", val)
	}
	if val := q.Pop(); val != 30 {
		t.Errorf("Pop() = %v, want 30", val)
	}
}

func TestQueue_FIFO(t *testing.T) {
	q := New[int](10)

	q.Push(1)
	q.Push(2)
	q.Push(3)

	if val := q.Pop(); val != 1 {
		t.Errorf("First pop should be 1, got %v", val)
	}
	if val := q.Pop(); val != 2 {
		t.Errorf("Second pop should be 2, got %v", val)
	}
	if val := q.Pop(); val != 3 {
		t.Errorf("Third pop should be 3, got %v", val)
	}
}

func TestQueue_IsEmpty(t *testing.T) {
	q := New[int](10)

	if !q.IsEmpty() {
		t.Error("New queue should be empty")
	}

	q.Push(42)
	if q.IsEmpty() {
		t.Error("Queue should not be empty after push")
	}

	q.Pop()
	if !q.IsEmpty() {
		t.Error("Queue should be empty after pop")
	}
}

func TestQueue_Size(t *testing.T) {
	q := New[int](10)

	if q.Size() != 0 {
		t.Errorf("Size() = %v, want 0", q.Size())
	}

	q.Push(1)
	q.Push(2)
	q.Push(3)

	if q.Size() != 3 {
		t.Errorf("Size() = %v, want 3", q.Size())
	}

	q.Pop()
	if q.Size() != 2 {
		t.Errorf("Size() after pop = %v, want 2", q.Size())
	}

	q.Pop()
	q.Pop()
	if q.Size() != 0 {
		t.Errorf("Size() after all pops = %v, want 0", q.Size())
	}
}
