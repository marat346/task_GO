package main

type Storage interface {
	Add(int) bool
	Size() int
	Values() []int
}

// затычка
type StubStorage struct{}

func (fs *StubStorage) Add(int) bool {
	return true
}
