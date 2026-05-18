package stack

import "testing"

func TestStack_PushAndPop(t *testing.T) {
	s := New[int](10)

	s.Push(10)
	s.Push(20)
	s.Push(30)

	if val := s.Pop(); val != 30 {
		t.Errorf("Pop() = %v, want 30", val)
	}
	if val := s.Pop(); val != 20 {
		t.Errorf("Pop() = %v, want 20", val)
	}
	if val := s.Pop(); val != 10 {
		t.Errorf("Pop() = %v, want 10", val)
	}
}

func TestStack_Peek(t *testing.T) {
	s := New[int](10)

	s.Push(100)
	s.Push(200)

	if val := s.Peek(); val != 200 {
		t.Errorf("Peek() = %v, want 200", val)
	}

	s.Pop()
	if val := s.Peek(); val != 100 {
		t.Errorf("Peek() after pop = %v, want 100", val)
	}
}

func TestStack_IsEmpty(t *testing.T) {
	s := New[int](10)

	if !s.IsEmpty() {
		t.Error("New stack should be empty")
	}

	s.Push(42)
	if s.IsEmpty() {
		t.Error("Stack should not be empty after push")
	}

	s.Pop()
	if !s.IsEmpty() {
		t.Error("Stack should be empty after pop")
	}
}

func TestStack_Size(t *testing.T) {
	s := New[int](10)

	if s.Size() != 0 {
		t.Errorf("Size() = %v, want 0", s.Size())
	}

	s.Push(1)
	s.Push(2)
	s.Push(3)

	if s.Size() != 3 {
		t.Errorf("Size() = %v, want 3", s.Size())
	}

	s.Pop()
	if s.Size() != 2 {
		t.Errorf("Size() after pop = %v, want 2", s.Size())
	}
}
