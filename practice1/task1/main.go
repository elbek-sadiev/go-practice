package main

import "fmt"

func bubbleSort(numbers []int) {
	for i := 0; i < len(numbers)-1; i++ {
		for j := 0; j < len(numbers)-i-1; j++ {
			if numbers[j] > numbers[j+1] {
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
			}
		}
	}
}

func main() {
	numbers := []int{42, 7, 19, 3, 88, 12, 1, 54}

	fmt.Println("Массив до сортировки:   ", numbers)
	bubbleSort(numbers)
	fmt.Println("Массив после сортировки:", numbers)
}
