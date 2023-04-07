package main

import "sync"

type server struct {
	status string
	sync.Mutex
}

func main() {

	s := server{}

	for i := 0; i < 1000; i++ {

		go s.Alive()
		go s.Down()
	}
}

func (s *server) Alive() {
	// Когда Mutex не разблок.Другая горутина не получит доступ к нашей структуре
	s.Lock()

	s.status = "Alive"

	s.Unlock()
}

func (s *server) Down() {
	s.Lock()

	s.status = "Down"

	s.Unlock()
}
