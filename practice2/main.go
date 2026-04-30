package main

import (
	"fmt"
	"math/rand"
)

type Stack struct {
	items []int
}

func (s *Stack) Push(value int) {
	s.items = append(s.items, value)
}

func (s *Stack) Pop() (int, bool) {
	if len(s.items) == 0 {
		return 0, false
	}

	lastIndex := len(s.items) - 1
	value := s.items[lastIndex]
	s.items = s.items[:lastIndex]

	return value, true
}

func (s *Stack) Peek() (int, bool) {
	if len(s.items) == 0 {
		return 0, false
	}

	return s.items[len(s.items)-1], true
}

type Queue struct {
	items []int
}

func (q *Queue) Push(value int) {
	q.items = append(q.items, value)
}

func (q *Queue) Pop() (int, bool) {
	if len(q.items) == 0 {
		return 0, false
	}

	value := q.items[0]
	q.items = q.items[1:]

	return value, true
}

type ListNode struct {
	value int
	next  *ListNode
}

type SinglyLinkedList struct {
	first *ListNode
	last  *ListNode
	size  int
}

func (l *SinglyLinkedList) Add(value int) {
	node := &ListNode{value: value}

	if l.first == nil {
		l.first = node
		l.last = node
	} else {
		l.last.next = node
		l.last = node
	}

	l.size++
}

func (l *SinglyLinkedList) Get(index int) (int, bool) {
	if index < 0 || index >= l.size {
		return 0, false
	}

	current := l.first

	for i := 0; i < index; i++ {
		current = current.next
	}

	return current.value, true
}

func (l *SinglyLinkedList) Remove(index int) bool {
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

	previous := l.first

	for i := 0; i < index-1; i++ {
		previous = previous.next
	}

	previous.next = previous.next.next

	if index == l.size-1 {
		l.last = previous
	}

	l.size--

	return true
}

func (l *SinglyLinkedList) Values() []int {
	values := make([]int, 0, l.size)
	current := l.first

	for current != nil {
		values = append(values, current.value)
		current = current.next
	}

	return values
}

func romanToArabic(roman string) int {
	values := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	result := 0

	for i := 0; i < len(roman); i++ {
		current := values[roman[i]]

		if i+1 < len(roman) && current < values[roman[i+1]] {
			result -= current
		} else {
			result += current
		}
	}

	return result
}

func createUniqueMatrix(rows, columns int) [][]int {
	total := rows * columns
	numbers := rand.Perm(total)

	matrix := make([][]int, rows)

	for i := 0; i < rows; i++ {
		matrix[i] = make([]int, columns)

		for j := 0; j < columns; j++ {
			matrix[i][j] = numbers[i*columns+j] + 1
		}
	}

	return matrix
}

func printMatrix(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Printf("%3d ", matrix[i][j])
		}

		fmt.Println()
	}
}

func main() {
	fmt.Println("Задание 1. Структуры данных")

	stack := Stack{}
	stack.Push(10)
	stack.Push(20)
	stack.Push(30)

	stackValue, _ := stack.Pop()
	fmt.Println("Стек, удаленный элемент:", stackValue)

	queue := Queue{}
	queue.Push(100)
	queue.Push(200)
	queue.Push(300)

	queueValue, _ := queue.Pop()
	fmt.Println("Очередь, удаленный элемент:", queueValue)

	list := SinglyLinkedList{}
	list.Add(5)
	list.Add(15)
	list.Add(25)
	list.Remove(1)

	fmt.Println("Односвязанный список:", list.Values())

	fmt.Println()
	fmt.Println("Задание 2. Конвертер римских цифр")

	roman := "MCMXCIV"
	arabic := romanToArabic(roman)

	fmt.Println("Римское число:", roman)
	fmt.Println("Арабское число:", arabic)

	fmt.Println()
	fmt.Println("Задание 3. Двумерный массив со случайными уникальными числами")

	matrix := createUniqueMatrix(4, 4)
	printMatrix(matrix)
}
