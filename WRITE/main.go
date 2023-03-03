package main

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
	a.PrintInfo
	a.PrintNumbers
	if number, ok := a.InputIndexNumbers(); ok {
		a.StoreNumbers(number)
	}
}

func (a *App) PrintInfo() {

}

func (a *App) PrintNumbers() {

}

func (a *App) InputIndexNumbers() {

}

func (a *App) StoreNumbers() {

}
