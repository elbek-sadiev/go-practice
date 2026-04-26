package main

import "fmt"

const maxEmployees = 512

type Employee struct {
	ID         int
	FirstName  string
	LastName   string
	Department string
	Position   string
}

type Office struct {
	employees []Employee
	nextID    int
}

func NewOffice() *Office {
	return &Office{
		employees: make([]Employee, 0, maxEmployees),
		nextID:    1,
	}
}

func (office *Office) AddEmployee(firstName, lastName, department, position string) bool {
	if len(office.employees) >= maxEmployees {
		return false
	}

	employee := Employee{
		ID:         office.nextID,
		FirstName:  firstName,
		LastName:   lastName,
		Department: department,
		Position:   position,
	}

	office.employees = append(office.employees, employee)
	office.nextID++

	return true
}

func (office *Office) DeleteEmployeeByID(id int) bool {
	for index, employee := range office.employees {
		if employee.ID == id {
			office.employees = append(office.employees[:index], office.employees[index+1:]...)
			return true
		}
	}

	return false
}

func (office *Office) PrintEmployees() {
	if len(office.employees) == 0 {
		fmt.Println("Список сотрудников пуст")
		return
	}

	for _, employee := range office.employees {
		fmt.Printf(
			"%d. %s %s, отдел: %s, должность: %s\n",
			employee.ID,
			employee.FirstName,
			employee.LastName,
			employee.Department,
			employee.Position,
		)
	}
}

func main() {
	office := NewOffice()

	office.AddEmployee("Иван", "Иванов", "Разработка", "Go-разработчик")
	office.AddEmployee("Анна", "Петрова", "Тестирование", "QA-инженер")
	office.AddEmployee("Павел", "Сидоров", "Аналитика", "Бизнес-аналитик")

	fmt.Println("Сотрудники после добавления:")
	office.PrintEmployees()

	deleted := office.DeleteEmployeeByID(2)
	if deleted {
		fmt.Println("\nСотрудник с ID 2 удален")
	} else {
		fmt.Println("\nСотрудник с ID 2 не найден")
	}

	fmt.Println("\nСотрудники после удаления:")
	office.PrintEmployees()
}
