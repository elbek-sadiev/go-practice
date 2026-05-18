package linkedlist

import (
	"reflect"
	"testing"
)

func TestLinkedList_AddAndValues(t *testing.T) {
	l := New[string]()

	l.Add("A")
	l.Add("B")
	l.Add("C")

	expected := []string{"A", "B", "C"}
	if !reflect.DeepEqual(l.Values(), expected) {
		t.Errorf("Values() = %v, want %v", l.Values(), expected)
	}
}

func TestLinkedList_Get(t *testing.T) {
	l := New[string]()

	l.Add("A")
	l.Add("B")
	l.Add("C")

	tests := []struct {
		idx  int
		want string
	}{
		{0, "A"},
		{1, "B"},
		{2, "C"},
	}

	for _, tt := range tests {
		if got := l.Get(tt.idx); got != tt.want {
			t.Errorf("Get(%d) = %v, want %v", tt.idx, got, tt.want)
		}
	}
}

func TestLinkedList_Remove(t *testing.T) {
	l := New[int]()

	l.Add(10)
	l.Add(20)
	l.Add(30)
	l.Add(40)

	l.Remove(1)

	expected := []int{10, 30, 40}
	if !reflect.DeepEqual(l.Values(), expected) {
		t.Errorf("After remove, Values() = %v, want %v", l.Values(), expected)
	}

	l.Remove(0)
	expected = []int{30, 40}
	if !reflect.DeepEqual(l.Values(), expected) {
		t.Errorf("After remove first, Values() = %v, want %v", l.Values(), expected)
	}

	l.Remove(1)
	expected = []int{30}
	if !reflect.DeepEqual(l.Values(), expected) {
		t.Errorf("After remove last, Values() = %v, want %v", l.Values(), expected)
	}
}

func TestLinkedList_Size(t *testing.T) {
	l := New[int]()

	if l.Size() != 0 {
		t.Errorf("Size() = %v, want 0", l.Size())
	}

	l.Add(1)
	l.Add(2)
	l.Add(3)

	if l.Size() != 3 {
		t.Errorf("Size() = %v, want 3", l.Size())
	}

	l.Remove(1)
	if l.Size() != 2 {
		t.Errorf("Size() after remove = %v, want 2", l.Size())
	}
}
