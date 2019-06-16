package d2

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"time"

	"github.com/nokka/slash-launcher/config"
	"github.com/nokka/slash-launcher/github"
	"github.com/nokka/slash-launcher/log"
)

// Service is responsible for all things related to Diablo II.
type Service struct {
	githubClient  github.Client
	configService config.Service
	logger        log.Logger
}

// Exec will exec the Diablo 2.
func (s *Service) Exec() error {
	conf, err := s.configService.Read()
	if err != nil {
		return err
	}

	cmd := exec.Command("cmd.exe", "/C", "call", `Diablo II.exe`, "-w")
	cmd.Dir = conf.D2Location

	err = cmd.Run()
	if err != nil {
		s.logger.Log(
			"location", conf.D2Location,
			"error", err,
		)
	}

	return nil
}

// Patch will check for updates and if found, patch the game.
func (s *Service) Patch(done chan bool) (<-chan float32, <-chan error) {
	progress := make(chan float32)
	errors := make(chan error)

	go func() {
		conf, err := s.configService.Read()
		if err != nil {
			s.logger.Log("msg", "failed to read config", "err", err)
			errors <- err
			return
		}

		// Download manifest from patch repository.
		manifest, err := s.getManifest()
		if err != nil {
			s.logger.Log("msg", "failed to get manifest", "err", err)
			errors <- err
			return
		}

		// Figure out which files to patch.
		patchFiles, patchLength, err := s.getFilesToPatch(manifest.Files, conf.D2Location)
		if err != nil {
			s.logger.Log("msg", "failed to get files to patch", "err", err)
			errors <- err
			return
		}

		// No files to patch, return.
		if len(patchFiles) == 0 {
			fmt.Println("NO FILES TO PATCH")
			return
		}

		// Create a write counter that will get bytes written per cycle, pass the
		// progress channel to report the number of bytes written.
		counter := &WriteCounter{
			Total:    float32(patchLength),
			progress: progress,
		}

		// Patch the files.
		for _, fileName := range patchFiles {
			f := fileName

			// Create the file, but give it a tmp file extension, this means we won't overwrite a
			// file until it's downloaded, but we'll remove the tmp extension once downloaded.
			tmp := fmt.Sprintf("%s/%s.tmp", conf.D2Location, f)

			out, err := os.Create(tmp)
			if err != nil {
				s.logger.Log("msg", "failed to create tmp file", "err", err)
				errors <- err
				return
			}

			defer out.Close()

			contents, err := s.githubClient.GetFile(f)
			if err != nil {
				s.logger.Log("failed to get file from github", err)
				errors <- err
				return
			}

			_, err = io.Copy(out, io.TeeReader(contents, counter))
			if err != nil {
				s.logger.Log("failed to write file locally", err)
				errors <- err
				return
			}

			// Download complete, remove the .tmp suffix.
			err = os.Rename(tmp, fmt.Sprintf("%s/%s", conf.D2Location, f))
			if err != nil {
				s.logger.Log("msg", "failed to rename tmp file", "err", err)
				errors <- err
				return
			}
		}

		done <- true
	}()

	return progress, errors
}

func (s *Service) getFilesToPatch(files []PatchFile, d2path string) ([]string, int64, error) {
	shouldPatch := make([]string, 0)
	var totalContentLength int64

	for _, file := range files {
		f := file

		// Full path on disk to the patch file.
		path := fmt.Sprintf("%s/%s", d2path, f.Name)

		// Get the checksum from the patch file on disk.
		hashed, err := hashCRC32(path, polynomial)
		if err != nil {
			// If the file doesn't exist on disk, we need to patch it.
			if err == ErrCRCFileNotFound {
				shouldPatch = append(shouldPatch, f.Name)
				totalContentLength += f.ContentLength
				continue
			}

			return nil, 0, err
		}

		// File checksum differs from local copy, we need to get a new one.
		if hashed != f.CRC {
			shouldPatch = append(shouldPatch, f.Name)
			totalContentLength += f.ContentLength
		}
	}

	return shouldPatch, totalContentLength, nil
}

func (s *Service) getManifest() (*Manifest, error) {
	contents, err := s.githubClient.GetFile("manifest.json")
	if err != nil {
		s.logger.Log("failed to get manifest from github", err)
		return nil, err
	}

	bytes, err := ioutil.ReadAll(contents)
	if err != nil {
		return nil, err
	}

	var manifest Manifest
	if err := json.Unmarshal(bytes, &manifest); err != nil {
		return nil, err
	}

	return &manifest, nil
}

// Manifest represents the current patch.
type Manifest struct {
	Files []PatchFile `json:"files"`
}

// PatchFile represents a file that should be patched.
type PatchFile struct {
	Name          string    `json:"name"`
	CRC           string    `json:"crc"`
	LastModified  time.Time `json:"last_modified"`
	ContentLength int64     `json:"content_length"`
}

// NewService returns a service with all the dependencies.
func NewService(
	githubClient github.Client,
	configuration config.Service,
	logger log.Logger,
) *Service {
	return &Service{
		githubClient:  githubClient,
		configService: configuration,
		logger:        logger,
	}
}
