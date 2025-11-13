package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/fatih/color"
)

func main() {
	fmt.Println("Hello, world!")

	// --- Слайд 5: Форматування коду ---
	x := 10
	if x > 0 {
		fmt.Println("--- Слайд 5: x is Positive ---")
	}

	fmt.Println(calculate(5, 3))

	var index int             // Погано, краще коротше
	var idx int               // Добре
	var requestHandler func() // Погано
	var h http.Handler        // Добре (коротше, контекст задає пакет)

	fmt.Println(index, idx, requestHandler, h)
	// --- Слайд 9: Коментарі (Добре) ---

	fmt.Println(factorial(5)) // Використання "гарної" версії
	fmt.Println("-----------------------------")

	go sayHello() // запуск паралельної goroutine

	ch := make(chan string)
	go func() {
		ch <- "Message from goroutine" // надсилаємо дані в канал
	}()
	msg := <-ch
	fmt.Println(msg) // Друкуємо повідомлення з каналу

	time.Sleep(1 * time.Second) // Даємо час горутині sayHello виконатись
	fmt.Println("Main function finished")
	color.Red("Text")
	ints := []int{1, 2, 3}
	PrintSlice(ints)
}

func calculate(a, b int) int {
	if a > b {
		return a + b
	} else {
		return b - a
	}
}

// Публічна структура (експортується)
type User struct {
	Name string // Публічне поле
	age  int    // Приватне поле (з малої літери)
}

// GetUserInfo - Публічна функція
func GetUserInfo() {}

// logMessage - Приватна функція
func logMessage() {}

// Добре: Використовуємо рекурсію для обходу дерева
func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

// Погано:
// Оголошуємо функцію факторіала (дублює код)
func factorial_bad(n int) int {
	// Якщо n менше або дорівнює 1, повертаємо 1 (очевидно)
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

// Це функція, яку ми будемо тестувати у `main_test.go`
func Add(a, b int) int {
	return a + b
}

// --- Слайд 13: Паралелізм (Goroutines) ---
func sayHello() {
	fmt.Println("Hello from goroutine!")
}

// --- Слайд 14: Загальні приклади (Гарний код з Generics)
func PrintSlice[T any](s []T) {
	for _, v := range s {
		fmt.Println(v)
	}
}
