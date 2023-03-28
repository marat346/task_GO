package easyLog

import (
	"fmt"
	"log"
	"os"
	"time"
)

func runAndWaitLog() int {
	time.Sleep(time.Second * 1)
	return 10
}

func main() {
	// создаем файл сразу,для записи логов
	// Флаги os.O_CREATE(Создание) или os.O_WRONLY(только писать) или os.O_APPEND(Добавить файл)

	file, err := os.OpenFile("test.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error while opening a file:%v\n", err)
	}

	// Запись в файл,чтоб не переделывает fmt в цикле,сразу в файл.
	log.SetOutput(file)

	for i := 0; i < 5; i++ {
		a := runAndWaitLog()
		// пишет сразу и время
		log.Println("runAndWait finished...")
		log.Println("result", a)
	}
	fmt.Println("done")
}
