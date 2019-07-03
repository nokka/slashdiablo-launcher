package d2

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
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

	if conf.D2Instances > 0 {
		for i := 0; i < conf.D2Instances; i++ {
			// Stall between each exec, otherwise Diablo won't start properly in multiple instances.
			time.Sleep(500 * time.Millisecond)
			go func() {
				if err := Exec(conf.D2Location); err != nil {
					s.logger.Log("msg", "failed to exec Diablo II", "err", err)
					return
				}
			}()
		}
	}

	if conf.HDInstances > 0 {
		for i := 0; i < conf.HDInstances; i++ {
			// Stall between each exec, otherwise Diablo won't start properly in multiple instances.
			time.Sleep(500 * time.Millisecond)
			go func() {
				if err := Exec(conf.HDLocation); err != nil {
					s.logger.Log("msg", "failed to exec HD Diablo II", "err", err)
					return
				}
			}()
		}
	}

	return nil
}

// Patch will check for updates and if found, patch the game, both D2 and HD version.
func (s *Service) Patch(done chan bool) (<-chan float32, <-chan PatchState) {
	progress := make(chan float32)
	state := make(chan PatchState)

	go func() {
		conf, err := s.configService.Read()
		if err != nil {
			s.logger.Log("msg", "failed to read config", "err", err)
			state <- PatchState{Error: err}
			return
		}

		// Check versions for both D2 and HD version.
		correctD2Version := validate113cVersion(conf.D2Location)
		correctHDVersion := validate113cVersion(conf.HDLocation)

		var pathsToUpdate []string
		if conf.D2Location != "" && !correctD2Version {
			pathsToUpdate = append(pathsToUpdate, conf.D2Location)
		}

		if conf.HDLocation != "" && !correctHDVersion {
			pathsToUpdate = append(pathsToUpdate, conf.HDLocation)
		}

		// Update Diablo installs to 1.13c if any of them are on the wrong version.
		if len(pathsToUpdate) > 0 {
			state <- PatchState{Message: "Updating Diablo II to 1.13c"}
			if err := s.apply113c(pathsToUpdate, progress); err != nil {
				s.logger.Log("msg", "failed to apply 1.13c", "err", err)
				state <- PatchState{Error: err}
				return
			}
		}

		state <- PatchState{Message: "Applying Slashdiablo patch"}

		// Apply latest Slashdiablo patch.
		if conf.D2Location != "" {
			err = s.applySlashPatch(conf.D2Location, progress)
			if err != nil {
				s.logger.Log("msg", "failed to patch slashdiablo patch", "err", err)
				state <- PatchState{Error: err}
				return
			}
		}

		// Apply patch to HD install if has been set.
		if conf.HDLocation != "" {
			err = s.applySlashPatch(conf.HDLocation, progress)
			if err != nil {
				s.logger.Log("msg", "failed to patch slashdiablo patch", "err", err)
				state <- PatchState{Error: err}
				return
			}
		}

		done <- true
	}()

	return progress, state
}

func (s *Service) apply113c(paths []string, progress chan float32) error {
	// Download manifest from patch repository.
	manifest, err := s.getManifest("1.13c/manifest.json")
	if err != nil {
		return err
	}

	for _, path := range paths {
		if err := s.doPatch(manifest.Files, path, progress); err != nil {
			// Make sure we clean up the failed patch.
			if err := s.cleanUpFailedPatch(path); err != nil {
				return err
			}

			return err
		}
	}

	return nil
}

func (s *Service) applySlashPatch(path string, progress chan float32) error {
	// Download manifest from patch repository.
	manifest, err := s.getManifest("current/manifest.json")
	if err != nil {
		return err
	}

	if err = s.doPatch(manifest.Files, path, progress); err != nil {
		// Make sure we clean up the failed patch.
		if err := s.cleanUpFailedPatch(path); err != nil {
			return err
		}

		return err
	}

	return nil
}

func (s *Service) doPatch(files []PatchFile, path string, progress chan float32) error {
	// Figure out which files to patch.
	patchFiles, patchLength, err := s.getFilesToPatch(files, path)
	if err != nil {
		return err
	}

	// Reset progress.
	progress <- 0.00

	// Create a write counter that will get bytes written per cycle, pass the
	// progress channel to report the number of bytes written.
	counter := &WriteCounter{
		Total:    float32(patchLength),
		progress: progress,
	}

	// Store the downloaded .tmp suffixed files.
	var tmpFiles []string

	// Patch the files.
	for _, fileName := range patchFiles {
		// Create the file, but give it a tmp file extension, this means we won't overwrite a
		// file until it's downloaded, but we'll remove the tmp extension once downloaded.
		tmpPath := localizePath(fmt.Sprintf("%s/%s.tmp", path, fileName))

		err := s.downloadFile(fileName, tmpPath, counter)
		if err != nil {
			return err
		}

		tmpFiles = append(tmpFiles, tmpPath)
	}

	// All the files were successfully downloaded, remove the .tmp suffix
	// to complete the patch entirely.
	for _, tmpFile := range tmpFiles {
		err = os.Rename(tmpFile, tmpFile[:len(tmpFile)-4])
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *Service) cleanUpFailedPatch(dir string) error {
	files, err := ioutil.ReadDir(localizePath(dir))
	if err != nil {
		return err
	}

	for _, f := range files {
		fileName := f.Name()
		if strings.Contains(fileName, ".tmp") {
			err := os.Remove(localizePath(fmt.Sprintf("%s/%s", dir, fileName)))
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (s *Service) downloadFile(fileName string, path string, counter *WriteCounter) error {
	out, err := os.Create(path)
	if err != nil {
		s.logger.Log("msg", "failed to create tmp file", "err", err)
		return err
	}

	defer out.Close()

	f := fmt.Sprintf("%s/%s", "1.13c", fileName)
	contents, err := s.githubClient.GetFile(f)
	if err != nil {
		s.logger.Log("failed to get file from github", err)
		return err
	}

	_, err = io.Copy(out, io.TeeReader(contents, counter))
	if err != nil {
		s.logger.Log("failed to write file locally", err)
		return err
	}

	return nil
}

func (s *Service) getFilesToPatch(files []PatchFile, d2path string) ([]string, int64, error) {
	shouldPatch := make([]string, 0)
	var totalContentLength int64

	for _, file := range files {
		f := file

		// Full path on disk to the patch file.
		path := localizePath(fmt.Sprintf("%s/%s", d2path, f.Name))

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

func (s *Service) getManifest(path string) (*Manifest, error) {
	contents, err := s.githubClient.GetFile(path)
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

// PatchState represents the state given on every patch cycle.
type PatchState struct {
	Message string
	Error   error
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
