package main

import (
	"fmt"
	"math/rand"
	"practice3/queue"
	"practice3/singlylist"
	"practice3/stack"
)

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

func createUniqueMatrix(rows, cols int) [][]int {
	numbers := rand.Perm(rows * cols)
	matrix := make([][]int, rows)

	for i := 0; i < rows; i++ {
		matrix[i] = make([]int, cols)
		for j := 0; j < cols; j++ {
			matrix[i][j] = numbers[i*cols+j] + 1
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

	s := stack.Stack[int]{}
	s.Push(10)
	s.Push(20)
	s.Push(30)
	stackValue, _ := s.Pop()
	fmt.Println("Стек, удаленный элемент:", stackValue)

	q := queue.Queue[int]{}
	q.Push(100)
	q.Push(200)
	q.Push(300)
	queueValue, _ := q.Pop()
	fmt.Println("Очередь, удаленный элемент:", queueValue)

	list := singlylist.List[int]{}
	list.Add(5)
	list.Add(15)
	list.Add(25)
	list.Remove(1)
	fmt.Println("Односвязанный список:", list.Values())

	fmt.Println()
	fmt.Println("Задание 2. Конвертер римских цифр")
	roman := "MCMXCIV"
	fmt.Println("Римское число:", roman)
	fmt.Println("Арабское число:", romanToArabic(roman))

	fmt.Println()
	fmt.Println("Задание 3. Двумерный массив со случайными уникальными числами")
	matrix := createUniqueMatrix(4, 4)
	printMatrix(matrix)
}
