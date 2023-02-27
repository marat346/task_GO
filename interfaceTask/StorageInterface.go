package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Storage interface {
	// Add добавляет число в коллекцию.
	// Возвращает true, если элемент добавлен;
	// false, если элемент уже присутствует.
	Add(int) bool

	// Size возвращает число элементов в коллекции.
	Size() int

	// Values возвращает срез с числами коллекции.
	Values() []int
}

// Создание «заглушки» репозитория
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

// Для этого создадим структуру App с единственным полем, которое хранит ссылку на репозиторий.
type App struct {
	repository Storage
}

// Теперь поэтапно создадим все методы структуры App. Начнём с главного цикла программы.
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

// Метод вывода информации о репозитории:
func (a *App) printInfo() {
	msg := "Уникальных чисел в колекции: %v\n"
	count := a.repository.Size()
	fmt.Printf(msg, count)
}

// Метод вывода всех чисел, введённых пользователем:
func (a *App) printNumbers() {
	fmt.Println("Список введенных значений:")
	fmt.Println(strings.Trim(fmt.Sprint(a.repository.Values()), "[]"))
}

// Метод ввода пользователем нового значения:
// Метод возвращает число и true, если введено корректное число или ноль, и false, если пользователь ввёл команду end.
func (a *App) inputNextNumber() (int, bool) {
	for {
		fmt.Print("Введите цифру или ~end~ для завершения:")
		var input string
		fmt.Scanln(&input)
		if value, err := strconv.Atoi(input); err == nil {
			return value, true
		}
		if input == "end" {
			return 0, false
		}
		fmt.Println("Некорректный ввод")
	}
}

// Метод, который добавляет новое число в репозиторий:
func (a *App) storeNumber(number int) {
	msg := "Число %d присутствует в коллекции\n"
	if ok := a.repository.Add(number); ok {
		msg = "Число %d успешно добавлено\n"
	}
	fmt.Printf(msg, number)
}

// Всё готово. Пришло время запустить приложение.
// Для этого объявим функцию main, в которой создадим экземпляры структур и передадим зависимости, то есть ссылки.
func main() {
	repository := &StubStorage{}
	app := &App{repository}
	app.Run()
} ///Обратите внимание: мы создали экземпляр структуры «заглушки»,
// а затем использовали ссылку на неё,
//чтобы создать экземпляр структуры App. Если всё сделано правильно, не должно быть ошибок компиляции,
//так как структура StubStorage имеет все методы, описанные в интерфейсе Storage.
