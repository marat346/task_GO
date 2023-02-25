package person

type Person struct {
	Age  int
	Name string
}

// Метод Ресивет по указателю
func (p *Person) SetNameByPointer(n string) {
	p.Name = n
}

// Метод Ресивет по копии
func (p Person) SetNameByValue(n string) {
	p.Name = n
}
