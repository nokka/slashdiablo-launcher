package d2

import (
	"fmt"

	"github.com/nokka/slash-launcher/github"
)

// Service is responible for all things related to Diablo II.
type Service struct {
	path          string
	githubService github.Service
}

// Exec will exec the Diablo 2.
func (s *Service) Exec() {
	fmt.Println("LAUNCH on path", s.path)
}

// Patch will check for updates and if found, patch the game.
func (s *Service) Patch() <-chan float32 {
	progress := make(chan float32)
	/*go func() {
		for i := 0; i <= 10; i++ {
			time.Sleep(1 * time.Second)
			p := 0.1 * float32(i)
			fmt.Println("SETTING PROGRESS", p)
			progress <- p
		}
		close(progress)
	}()*/
	go func() {
		content, err := s.githubService.GetFile("README.md")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(content)
	}()

	return progress

}

// NewService returns a service with all the dependencies.
func NewService(path string, githubService github.Service) *Service {
	return &Service{
		path:          path,
		githubService: githubService,
	}
}
