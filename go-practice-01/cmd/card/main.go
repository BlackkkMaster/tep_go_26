package main

import "fmt"

func main() {
	const course = "Бэкенд-разработка на Go"
	group := "Учебная-03"
	name := "Тони Старк"
	var completed int = 2
	interesting := true

	fmt.Printf("Участник: %s\n", name)
	fmt.Printf("Группа: %s\n", group)
	fmt.Printf("Курс: %s\n", course)
	fmt.Printf("Завершено работ: %d\n", completed)
	fmt.Printf("Интерес к бэкенду: %t\n", interesting)
}
