package d2

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/nokka/slashdiablo-launcher/config"
	"github.com/nokka/slashdiablo-launcher/github"
)

// Service is responsible for all things related to Diablo II.
type Service struct {
	githubClient  github.Client
	configService config.Service
	gameStates    chan execState
	runningGames  []game
	mux           sync.Mutex
}

type game struct {
	PID    int
	GameID int
}

type execState struct {
	pid *int
	err error
}

func (s *Service) listenForGameStates() {
	fmt.Println("Listening for game states")
	for {
		select {
		case state := <-s.gameStates:
			fmt.Println("GOT STATE-------------------------")
			// Something went wrong while execing, log error.
			if state.err != nil {
				fmt.Println("Got error", state.err)
			}

			s.mux.Lock()

			// Game exited, remove it from the slice based on pid.
			for index, g := range s.runningGames {
				if state.pid != nil && g.PID == *state.pid {
					fmt.Println("Delete index", index)
					s.runningGames = append(s.runningGames[:index], s.runningGames[index+1:]...)

					fmt.Println("Running games after the delete")
					fmt.Println(s.runningGames)
				}
			}

			s.mux.Unlock()
			fmt.Println("------------------------------")
		}
	}
}

// Exec will exec the Diablo 2.
func (s *Service) Exec() error {
	conf, err := s.configService.Read()
	if err != nil {
		return err
	}

	for _, g := range conf.Games {
		for i := 0; i < g.Instances; i++ {
			// Stall between each exec, otherwise Diablo won't start properly in multiple instances.
			time.Sleep(200 * time.Millisecond)
			pid, err := launch(g.Location, s.gameStates)
			if err != nil {
				return err
			}

			// Add the started game to our slice of games.
			s.runningGames = append(s.runningGames, game{PID: *pid, GameID: g.ID})
		}
	}

	fmt.Println("RUNNING GAMES")
	fmt.Println(s.runningGames)

	return nil
}

// ValidateGameVersions will check if the games are up to date.
func (s *Service) ValidateGameVersions() (bool, error) {
	conf, err := s.configService.Read()
	if err != nil {
		return false, err
	}

	// Get current slash patch and compare.
	manifest, err := s.getManifest("current/manifest.json")
	if err != nil {
		return false, err
	}

	if len(conf.Games) > 0 {
		for _, game := range conf.Games {
			ok, err := validate113cVersion(game.Location)
			if err != nil {
				return false, err
			}

			// Game wasn't 1.13c, needs to be updated.
			if !ok {
				return false, nil
			}

			// Check if the current game install is up to date with the slash patch.
			patchFiles, _, err := s.getFilesToPatch(manifest.Files, game.Location)
			if err != nil {
				return false, err
			}

			if len(patchFiles) > 0 {
				return false, nil
			}

			// Game is both 1.13c and up to date with Slash patches.
			return true, nil
		}
	}

	return false, nil
}

// Patch will check for updates and if found, patch the game, both D2 and HD version.
func (s *Service) Patch(done chan bool) (<-chan float32, <-chan PatchState) {
	progress := make(chan float32)
	state := make(chan PatchState)

	go func() {
		conf, err := s.configService.Read()
		if err != nil {
			state <- PatchState{Error: err}
			return
		}

		for _, game := range conf.Games {
			if err := s.apply113c(game.Location, state, progress); err != nil {
				state <- PatchState{Error: err}
				return
			}

			err = s.applySlashPatch(game.Location, state, progress)
			if err != nil {
				state <- PatchState{Error: err}
				return
			}

			if game.Maphack {
				err = s.applyMaphack(game.Location, state, progress)
				if err != nil {
					state <- PatchState{Error: err}
					return
				}
			}

			if game.HD {
				err = s.applyHDMod(game.Location, state, progress)
				if err != nil {
					state <- PatchState{Error: err}
					return
				}
			}
		}

		done <- true
	}()

	return progress, state
}

// RunDEPFix will run a specific fix to disable DEP.
func (s *Service) RunDEPFix() error {
	/*conf, err := s.configService.Read()
	if err != nil {
		return err
	}

	// Run OS specific fix.
	err = runDEPFix("")
	if err != nil {
		return err
	}*/

	return nil
}

func (s *Service) apply113c(path string, state chan PatchState, progress chan float32) error {
	// Update UI.
	state <- PatchState{Message: "Validating game version"}

	// Download manifest from patch repository.
	manifest, err := s.getManifest("1.13c/manifest.json")
	if err != nil {
		return err
	}

	// Figure out which files to patch.
	patchFiles, patchLength, err := s.getFilesToPatch(manifest.Files, path)
	if err != nil {
		return err
	}

	if len(patchFiles) > 0 {
		fmt.Println("SHOULD PATCH")
		fmt.Println("------------------")
		fmt.Println(path)
		fmt.Println(patchFiles)
		fmt.Println("------------------")
		// Update UI.
		state <- PatchState{Message: fmt.Sprintf("Updating %s to 1.13c", path)}

		if err := s.doPatch(patchFiles, patchLength, "1.13c", path, progress); err != nil {
			// Make sure we clean up the failed patch.
			if err := s.cleanUpFailedPatch(path); err != nil {
				return err
			}

			return err
		}
	}

	return nil
}

func (s *Service) applySlashPatch(path string, state chan PatchState, progress chan float32) error {
	// Update UI.
	state <- PatchState{Message: "Validating Slashdiablo patch"}

	// Download manifest from patch repository.
	manifest, err := s.getManifest("current/manifest.json")
	if err != nil {
		return err
	}

	// Figure out which files to patch.
	patchFiles, patchLength, err := s.getFilesToPatch(manifest.Files, path)
	if err != nil {
		return err
	}

	if len(patchFiles) > 0 {
		// Update UI.
		state <- PatchState{Message: fmt.Sprintf("Updating %s to current Slashdiablo patch", path)}

		if err = s.doPatch(patchFiles, patchLength, "current", path, progress); err != nil {
			// Make sure we clean up the failed patch.
			if err := s.cleanUpFailedPatch(path); err != nil {
				return err
			}

			return err
		}
	}

	return nil
}

func (s *Service) applyMaphack(path string, state chan PatchState, progress chan float32) error {
	// Update UI.
	state <- PatchState{Message: "Validating maphack"}

	// Download manifest from patch repository.
	manifest, err := s.getManifest("maphack/manifest.json")
	if err != nil {
		return err
	}

	// Figure out which files to patch.
	patchFiles, patchLength, err := s.getFilesToPatch(manifest.Files, path)
	if err != nil {
		return err
	}

	if len(patchFiles) > 0 {
		// Update UI.
		state <- PatchState{Message: fmt.Sprintf("Updating %s to latest maphack version", path)}

		if err = s.doPatch(patchFiles, patchLength, "maphack", path, progress); err != nil {
			// Make sure we clean up the failed patch.
			if err := s.cleanUpFailedPatch(path); err != nil {
				return err
			}

			return err
		}
	}

	return nil
}

func (s *Service) applyHDMod(path string, state chan PatchState, progress chan float32) error {
	// Update UI.
	state <- PatchState{Message: "Validating HD mod"}

	// Download manifest from patch repository.
	manifest, err := s.getManifest("hd/manifest.json")
	if err != nil {
		return err
	}

	// Figure out which files to patch.
	patchFiles, patchLength, err := s.getFilesToPatch(manifest.Files, path)
	if err != nil {
		return err
	}

	if len(patchFiles) > 0 {
		// Update UI.
		state <- PatchState{Message: fmt.Sprintf("Updating %s to latest HD mod version", path)}

		if err = s.doPatch(patchFiles, patchLength, "hd", path, progress); err != nil {
			// Make sure we clean up the failed patch.
			if err := s.cleanUpFailedPatch(path); err != nil {
				return err
			}

			return err
		}
	}

	return nil
}

func (s *Service) doPatch(patchFiles []string, patchLength int64, remoteDir string, path string, progress chan float32) error {
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

		err := s.downloadFile(fileName, remoteDir, tmpPath, counter)
		if err != nil {
			return err
		}

		tmpFiles = append(tmpFiles, tmpPath)
	}

	// All the files were successfully downloaded, remove the .tmp suffix
	// to complete the patch entirely.
	for _, tmpFile := range tmpFiles {
		err := os.Rename(tmpFile, tmpFile[:len(tmpFile)-4])
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *Service) downloadFile(fileName string, remoteDir string, path string, counter *WriteCounter) error {
	out, err := os.Create(path)
	if err != nil {
		return err
	}

	defer out.Close()

	f := fmt.Sprintf("%s/%s", remoteDir, fileName)
	contents, err := s.githubClient.GetFile(f)
	if err != nil {
		return err
	}

	_, err = io.Copy(out, io.TeeReader(contents, counter))
	if err != nil {
		return err
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

		fmt.Println("------------------")
		fmt.Println("FILE", f.Name)
		fmt.Println("LOCAL", hashed)
		fmt.Println("REMOTE", f.CRC)

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
) *Service {
	s := &Service{
		githubClient:  githubClient,
		configService: configuration,
		gameStates:    make(chan execState, 4),
	}

	// Setup game listener once, will stay alive for the duration
	// of the service's life cycle.
	go s.listenForGameStates()

	return s
}
