package WRITE

type Storage interface {
	Add(int) bool
	Size() int
	Values() []int
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
