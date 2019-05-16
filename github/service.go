package github

import (
	"context"
	"io"
	"sync"

	"github.com/google/go-github/github"
)

// Service encapsulates all the operations available on the github service.
type Service interface {
	GetFile(filePath string) (io.ReadCloser, error)
}

type service struct {
	owner      string
	repository string
	mutex      sync.Mutex
	client     *github.Client
	ctx        context.Context
}

// GetFile will the file by the given path in the repository set on the service.
func (s *service) GetFile(filePath string) (io.ReadCloser, error) {
	client, err := s.getClient()
	if err != nil {
		return nil, err
	}
	return client.Repositories.DownloadContents(
		s.ctx,
		s.owner,
		s.repository,
		filePath,
		nil,
	)
}

// getClient creates the github API client if its not set already.
func (s *service) getClient() (*github.Client, error) {
	// Lock in case multiple threads are trying to get
	// the client at the same time.
	s.mutex.Lock()

	// Unlock when we're done mutating the client.
	defer s.mutex.Unlock()

	if s.client == nil {
		s.client = github.NewClient(nil)
	}

	return s.client, nil
}

// NewService returns a new github service with all dependencies setup.
func NewService(owner string, repository string) Service {
	return &service{
		owner:      owner,
		repository: repository,
		ctx:        context.Background(),
	}
}
