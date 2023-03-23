package main

import "fmt"

func main() {
	var sum int

	fmt.Println("Банкомат.")
	fmt.Println("Введите сумму снятия со счёта:")
	fmt.Scan(&sum)
	if sum%100 > 0 {
		fmt.Println("Банкомат выдаёт только купюры номиналом 100 рублей")
	} else if sum > 100000 {
		fmt.Println("Максимальная сумма снятия — 100 000 рублей")
	} else {
		fmt.Println("Операция успешно выполнена.")
		fmt.Println("Вы сняли", sum, "рублей.")
	}
}
