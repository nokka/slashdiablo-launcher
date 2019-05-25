package d2

import (
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/nokka/slash-launcher/config"
	"github.com/nokka/slash-launcher/github"
	"github.com/nokka/slash-launcher/log"
)

// Service is responsible for all things related to Diablo II.
type Service struct {
	githubService github.Service
	configService config.Service
	logger        log.Logger
}

// Exec will exec the Diablo 2.
func (s *Service) Exec() error {
	conf, err := s.configService.Read()
	if err != nil {
		return err
	}

	fmt.Println("LAUNCH ON LOCATION", conf.D2Location)
	fmt.Println("WITH INSTANCES", conf.D2Instances)

	command := "open"
	arg1 := "-a"
	arg2 := fmt.Sprintf("%s/Diablo II.exe", conf.D2Location)
	cmd := exec.Command(command, arg1, arg2)

	err = cmd.Run()
	if err != nil {
		s.logger.Log("unable to exec game path", conf.D2Location, err)
		s.logger.Log(errors.New("My error"))
		return err
	}

	return nil
}

// Patch will check for updates and if found, patch the game.
func (s *Service) Patch() <-chan float32 {
	progress := make(chan float32)

	go func() {
		defer close(progress)

		conf, err := s.configService.Read()
		if err != nil {
			s.logger.Log("failed to read config", err)
			return
		}

		// Create the file, but give it a tmp file extension, this means we won't overwrite a
		// file until it's downloaded, but we'll remove the tmp extension once downloaded.
		out, err := os.Create(conf.D2Location + "/test.tmp")
		if err != nil {
			s.logger.Log("failed to create tmp file", err)
			return
		}

		defer out.Close()

		contents, err := s.githubService.GetFile("Patch_D2.mpq")
		if err != nil {
			s.logger.Log("failed to get file from github", err)
			return
		}

		// Create our progress reporter and pass it to be used alongside our writer
		counter := &WriteCounter{
			Total:    float32(2108703),
			progress: progress,
		}

		_, err = io.Copy(out, io.TeeReader(contents, counter))
		if err != nil {
			s.logger.Log("failed to write file locally", err)
			return
		}
	}()

	return progress

}

// NewService returns a service with all the dependencies.
func NewService(
	githubService github.Service,
	configuration config.Service,
	logger log.Logger,
) *Service {
	return &Service{
		githubService: githubService,
		configService: configuration,
		logger:        logger,
	}
}
