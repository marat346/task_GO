package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"main/student"
)

func main() {
	fmt.Println("Введите имя студента возраст и курс через пробел")
	fmt.Println("Для завершения введите комбинацию клавишь Ctrl+D")
  
	mapStudent := make(map[string]*student.Student, 0)
	var in = bufio.NewReader(os.Stdin)

	for {
		str, err := in.ReadString('\n')

		if err == io.EOF {
			break
		}
		addNewStudent(mapStudent, str)
	}

	printStudent(mapStudent)
}

func addNewStudent(mapStudent map[string]*student.Student, str string) {
	name, age, grade := getContent(str)
  mapStudent[name] = student.Put(name, age, grade)
}

func getContent(str string) (name string, age int, grade int) {
	parseStr := strings.Fields(str)

	name = parseStr[0]
	age, _ = strconv.Atoi(parseStr[1])
	grade, _ = strconv.Atoi(parseStr[2])

	return
}

func printStudent(mapStudent map[string]*student.Student) {
  fmt.Println()
	fmt.Println("Студенты из хранилища:")
	for _, v := range mapStudent {
    student.Print(*v)
	}
}