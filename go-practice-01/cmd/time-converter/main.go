package main

import "fmt"

func main() {
	var minuts int
	const min_per_hour = 60
	fmt.Print("Введите количество минут: ")
	_, err := fmt.Scan(&minuts)
	if err != nil {
		fmt.Println("Ошибка: нужно ввести целое число.")
		return
	}
	if minuts < 0 {
		fmt.Println("Ошибка: число не должно быть отрицательным.")
		return
	}
	var hours int = minuts / min_per_hour
	var mins_left int = minuts % min_per_hour

	fmt.Printf("Полных часов: %d\n", hours)
	fmt.Printf("Оставшихся минут: %d\n", mins_left)
}
