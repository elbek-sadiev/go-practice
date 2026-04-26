package main

import "fmt"

type Employee struct {
	Name     string
	Age      int
	Position string
	Salary   int
}

const size = 512

var commands = `
1 - Добавить нового сотрудника
2 - Удалить сотрудника
3 - Вывести список сотрудников
4 - Выйти из программы
`

func main() {
	employees := [size]*Employee{}

	for {
		var command int

		fmt.Print(commands)
		fmt.Print("Введите команду: ")
		fmt.Scan(&command)

		switch command {
		case 1:
			addEmployee(&employees)
		case 2:
			deleteEmployee(&employees)
		case 3:
			printEmployees(employees)
		case 4:
			fmt.Println("Программа завершена")
			return
		default:
			fmt.Println("Неизвестная команда")
		}
	}
}

func addEmployee(employees *[size]*Employee) {
	index := -1

	for i := 0; i < size; i++ {
		if employees[i] == nil {
			index = i
			break
		}
	}

	if index == -1 {
		fmt.Println("Нельзя добавить больше 512 сотрудников")
		return
	}

	employee := new(Employee)

	fmt.Print("Имя: ")
	fmt.Scan(&employee.Name)

	fmt.Print("Возраст: ")
	fmt.Scan(&employee.Age)

	fmt.Print("Должность: ")
	fmt.Scan(&employee.Position)

	fmt.Print("Зарплата: ")
	fmt.Scan(&employee.Salary)

	employees[index] = employee

	fmt.Println("Сотрудник добавлен")
}

func deleteEmployee(employees *[size]*Employee) {
	var number int

	fmt.Print("Введите номер сотрудника для удаления: ")
	fmt.Scan(&number)

	index := number - 1

	if index < 0 || index >= size || employees[index] == nil {
		fmt.Println("Сотрудник не найден")
		return
	}

	employees[index] = nil

	fmt.Println("Сотрудник удален")
}

func printEmployees(employees [size]*Employee) {
	hasEmployees := false

	for i := 0; i < size; i++ {
		if employees[i] != nil {
			fmt.Printf(
				"%d. Имя: %s, возраст: %d, должность: %s, зарплата: %d\n",
				i+1,
				employees[i].Name,
				employees[i].Age,
				employees[i].Position,
				employees[i].Salary,
			)

			hasEmployees = true
		}
	}

	if !hasEmployees {
		fmt.Println("Список сотрудников пуст")
	}
}
