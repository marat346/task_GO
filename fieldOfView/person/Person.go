package person

import "strconv"

// так же задаем константу C,если она с мал.буквы main знать о ней не будет
const C = "from person package "

// есть способ сделать видимой константу(можно и поля и переменные) с помощью метода публичного метода GetString
const f = "сделайте меня видимой из person"

func GetString() string {
	return f
}

// Если не хотим экспортировать всю струтуру,ставим ее с мал.буквы person
type Person struct {
	Age      int
	Name     string
	Sibiling string
}

// Так же если функция написана с маленькой буквы,она неэкспортируема printInfo

//       func PrintInfo(p Person) string {
//         return "My name is " + p.Name + " and age is " + strconv.Itoa(p.Age) + p.Sibiling
//        }

// Функция в функции.Одна с большой буквы другая с маленькой
func PrintInfo(p Person) string {
	return info(p)
}

// код main (клиентский код) не хотел знать логику, из пакета person скрыть
// ставим функцию с маленькой буквы

func info(p Person) string {
	return "My name is " + p.Name + " and age is " + strconv.Itoa(p.Age) + p.Sibiling
}
