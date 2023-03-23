package main

import (
	"fmt"
	"sort"
)

func main() {
	//создаем map
	m := map[int]string{
		1: "first",
		2: "second",
		3: "third",
		4: "forth",
	}
	// беспорядочный цикл ,индекс меняется каждый раз при компиляции
	for k := range m {
		fmt.Println(k, "все ключи")
	}

	fmt.Println()
	// беспорядочный цикл ,индекс меняется каждый раз при компиляции
	for v, k := range m {
		fmt.Println(v, k)
	}

	// способ как упорядочить вывод цикла
	keySlice := make([]int, 0)
	for k := range m {
		keySlice = append(keySlice, k)
	}
	fmt.Println("___________")
	fmt.Println(keySlice)
	fmt.Println("________________")
	// сортируем слайс и через цикл выводим
	sortedKeys := sort.IntSlice(keySlice)
	fmt.Println("ЭТО СОРТИРОВАННЫЙ СЛАЙС", sortedKeys)
	fmt.Println("_______________________________")
	for _, key := range sortedKeys {
		fmt.Println(m[key])
	}

}
