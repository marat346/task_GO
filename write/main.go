package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Storage interface {
	Add(int) bool
	Size() int
	Values() []int
}

type StubStorage struct{}

func (fs *StubStorage) Add(int) bool {
	return true
}

func (fs *StubStorage) Size() int {
	return 3
}

func (fs *StubStorage) Values() []int {
	return []int{1, 2, 3}
}

type App struct {
	repository Storage
}

func (a *App) Run() {
	for {
		a.printInfo()
		a.printNumbers()
		if number, ok := a.inputNextNumber(); ok {
			a.storeNumber(number)
		} else {
			break
		}
	}
}

func (a *App) printInfo() {
	msg := "Уникальных чисел в коллекции : %v/n"
	count := a.repository.Size()
	fmt.Printf(msg, count)
}

func (a *App) printNumbers() {
	fmt.Println("Список введенных значений:")
	fmt.Println(strings.Trim(fmt.Sprint(a.repository.Values()), "[]"))
}

func (a *App) inputNextNumber() (int, bool) {
	for {
		fmt.Print("Введите цифру или `end` для завершения :")
		var input string
		fmt.Scan(&input)
		if value, err := strconv.Atoi(input); err == nil {
			return value, true
		}
		if input == "end" {
			return 0, false
		}
		fmt.Println("Некорректный ввод")
	}
}

func (a *App) storeNumber(number int) {
	msg := "Число %d присутствует в коллекции\n"
	if ok := a.repository.Add(number); ok {
		msg = "Число успешно добавлено \n"
	}
	fmt.Printf(msg, number)
}

func main() {
	repository := &StubStorage{}
	app := &App{repository}
	app.Run()
}
