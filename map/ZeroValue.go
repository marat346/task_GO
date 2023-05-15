package main

import "fmt"

func main() {
	mLen := map[int]string{
		1: "first",
		2: "second",
		3: "third",
	}

	fmt.Println(len(mLen))

	//ZeroValue получается ошибка
	fmt.Println(mLen[4])

	// можно воспользоваться синтаксисом
	//v попадет ошибка ZeroValue(пустая строка),а в Ok выйдет false(нету ключа) или true (ключ есть)
	v, ok := mLen[4]
	fmt.Println(v, ok)
	//пустая строка
	// false

	//чтоб не возращать значение ,а просто проверить исп.Получаем false
	_, ok = mLen[4]
	fmt.Println(ok)
	// Проверка на true,получаем true.Ключ есть
	_, ok = mLen[3]
	fmt.Println(ok)

}
