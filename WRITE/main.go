package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Storage interface {
	Add(int) bool
	Size() int
	Values() int
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
		a.PrintInfo
		a.PrintNumbers
		if number, ok := a.InputIndexNumbers(); ok {
			a.StoreNumbers(number)
		} else {
			break
		}
	}
}

func (a *App) PrintInfo() {
	msg := fmt.Println("Введите число")
	count := a.repository.Size()
	fmt.Println(msg, count)
}

func (a *App) PrintNumbers() {
fmt.Println("jljljlkjlkjjklj")
fmt.Println(strings.Trim(fmt.Sprint(a.repository.Values()),[]))
}

func (a *App) InputIndexNumbers() (int,bool){
for{
	fmt.Print("jjljljljlk")
	var input string
	fmt.Scanln(&input)
	if number , err := strconv.Atoi(input);err == nil{
		return value , true
	}
	if input == "end" {

	}
}
}

func (a *App) StoreNumbers() {

}
