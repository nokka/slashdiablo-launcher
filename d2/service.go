package d2

import (
	"fmt"
	"time"
)

// Service is responible for all things related to Diablo II.
type Service struct {
	Path string
}

// Exec will exec the Diablo 2.
func (s *Service) Exec() {
	fmt.Println("LAUNCH on path", s.Path)
}

// Patch will check for updates and if found, patch the game.
func (s *Service) Patch() <-chan float32 {
	progress := make(chan float32)
	go func() {
		for i := 0; i <= 10; i++ {
			time.Sleep(1 * time.Second)
			p := 0.1 * float32(i)
			fmt.Println("SETTING PROGRESS", p)
			progress <- p
		}
		close(progress)
	}()
	return progress

}

// NewService returns a service with all the dependencies.
func NewService(path string) *Service {
	return &Service{
		Path: path,
	}
}
