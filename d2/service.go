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
	"github.com/nokka/slashdiablo-launcher/log"
)

// Service is responsible for all things related to Diablo II.
type Service struct {
	githubClient  github.Client
	configService config.Service
	logger        log.Logger
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
	for {
		select {
		case state := <-s.gameStates:
			// Something went wrong while execing, log error.
			if state.err != nil {
				s.logger.Error(fmt.Errorf("Diablo II exec with code: %s", state.err))
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
		}
	}
}

// SetGateway will set the given gateway for the user.
func (s *Service) SetGateway(gateway string) error {
	return setGateway(gateway)
}

/*func mutateInstancesToLaunch([]storage.Game) {
	// Loop over config games

	// Loop over running instances

	// Count how many instances are running of the specific id

	// Subtract the running games from game.Instances

	// Add the new instances to launch int to the config game.
}*/

// Exec will exec Diablo 2 installs.
func (s *Service) Exec() error {
	conf, err := s.configService.Read()
	if err != nil {
		return err
	}

	// Mutate the number of instances to launch to take into
	// account the number of games already running.
	//mutateInstancesToLaunch(conf.Games)

	for _, g := range conf.Games {
		for i := 0; i < g.Instances; i++ {
			// Stall between each exec, otherwise Diablo won't start properly in multiple instances.
			time.Sleep(1 * time.Second)

			// The second argument is a channel, listened on by listenForGameStates().
			pid, err := launch(g.Location, s.gameStates)
			if err != nil {
				return err
			}

			// Add the started game to our slice of games.
			s.runningGames = append(s.runningGames, game{PID: *pid, GameID: g.ID})
		}
	}

	return nil
}

// ValidateGameVersions will check if the games are up to date.
func (s *Service) ValidateGameVersions() (bool, error) {
	conf, err := s.configService.Read()
	if err != nil {
		return false, err
	}

	// Get current slash patch and compare.
	slashManifest, err := s.getManifest("current/manifest.json")
	if err != nil {
		return false, err
	}

	// Get current maphack patch and compare.
	maphackManifest, err := s.getManifest("maphack/manifest.json")
	if err != nil {
		return false, err
	}

	fmt.Println("MAPHACK MANIFEST ---")
	fmt.Println(maphackManifest)

	// Get current HD patch and compare.
	HDManifest, err := s.getManifest("hd/manifest.json")
	if err != nil {
		return false, err
	}

	upToDate := true

	if len(conf.Games) > 0 {
		for _, game := range conf.Games {
			valid, err := validate113cVersion(game.Location)
			if err != nil {
				return false, err
			}

			// Game wasn't 1.13c, needs to be updated.
			if !valid {
				fmt.Println("1.13c WAS NOT UP TO DATE, RETURNING FALSE")
				return false, nil
			}

			fmt.Println("1.13c WAS UP TO DATE")

			// Check if the current game install is up to date with the slash patch.
			slashFiles, _, err := s.getFilesToPatch(slashManifest.Files, game.Location, nil)
			if err != nil {
				return false, err
			}

			// Slash patch isn't up to date.
			if len(slashFiles) > 0 {
				fmt.Println("SLASH PATCH IS NOT UP TO DATE")
				return false, nil
			}

			fmt.Println("SLASH PATCH WAS UP TO DATE")

			// If the user has chosen to override the maphack config with their own,
			// we need to make sure the config is being ignored from the patch.
			var ignoredMaphackFiles []string

			if game.OverrideBHCfg {
				ignoredMaphackFiles = append(ignoredMaphackFiles, "BH.cfg")
			}

			fmt.Println("IGNORED FILES WHILE VALIDATING VERSION", ignoredMaphackFiles)

			// Check how many files aren't up to date with maphack.
			missingMaphackFiles, _, err := s.getFilesToPatch(maphackManifest.Files, game.Location, ignoredMaphackFiles)
			if err != nil {
				return false, err
			}

			fmt.Println("MISSING MAPHACK FILES", missingMaphackFiles)

			// Maphack is enabled, make sure there's no missing files.
			if game.Maphack {
				// Maphack patch isn't up to date.
				if len(missingMaphackFiles) > 0 {
					fmt.Println("MAPHACK WAS ENABLED, BUT OUT OF DATE / MISSING FILES")
					fmt.Println(missingMaphackFiles)
					return false, nil
				}
				fmt.Println("MAPHACK WAS ENABLED, EVERYTHING WAS OK")
			} else {
				installed, err := isMaphackInstalled(game.Location)
				if err != nil {
					return false, err
				}

				// Maphack wasn't supposed to be installed, but it is, we need to update.
				if installed {
					fmt.Println("MAPHACK WAS INSTALLED WHEN IT WASN'T SUPPOSED TO, NEED TO UPDATE")
					return false, nil
				}

				fmt.Println("MAPHACK WAS NOT ENABLED, EVERYTHING WAS OK")
			}

			// Check if the current game install is up to date with the HD patch.
			missingHDFiles, _, err := s.getFilesToPatch(HDManifest.Files, game.Location, nil)
			if err != nil {
				return false, err
			}

			if game.HD {
				// HD patch isn't up to date.
				if len(missingHDFiles) > 0 {
					fmt.Println("HD WAS ENABLED, BUT MISSING FILES")
					return false, nil
				}
				fmt.Println("HD WAS ENABLED, EVERYTHING WAS OK")
			} else {
				installed, err := isHDInstalled(game.Location)
				if err != nil {
					return false, err
				}

				// HD wasn't supposed to be installed, but it is, we need to update.
				if installed {
					fmt.Println("HD WAS INSTALLED WHEN IT WASN'T SUPPOSED TO, NEED TO UPDATE")
					return false, nil
				}

				fmt.Println("HD WAS NOT ENABLED, EVERYTHING WAS OK")
			}
		}
	}

	// Games are both 1.13c and up to date with Slash patch, maphack and HD.
	return upToDate, nil
}

func (s *Service) resetPatch(path string, files []PatchFile, filesToIgnore []string) error {
	// Check how many files aren't up to date.
	missmatchedFiles, _, err := s.getFilesToPatch(files, path, filesToIgnore)
	if err != nil {
		return err
	}

	// If the number of missmatched files to patch aren't all of them, then we have
	// some of them left that needs to be removed.
	if len(missmatchedFiles) != len(files) {
		for _, file := range files {
			filePath := localizePath(fmt.Sprintf("%s/%s", path, file.Name))

			// Check if the file exists, on disk, if it does, remove it.
			_, err := os.Stat(filePath)
			if err != nil {
				// File didn't exist on disk, continue to next.
				if os.IsNotExist(err) {
					continue
				}
				// Unknown error.
				return err

			}

			// Make sure we don't remove the ignored files.
			var ignore bool

			for _, ignored := range filesToIgnore {
				if file.Name == ignored {
					ignore = true
					break
				}
			}

			if !ignore {
				// File that shouldn't be on disk exists, remove it.
				err = os.Remove(filePath)
				if err != nil {
					return err
				}
			}
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
			state <- PatchState{Error: err}
			return
		}

		// Download maphack manifest from patch repository, , we'll use it multiple times.
		maphackManifest, err := s.getManifest("maphack/manifest.json")
		if err != nil {
			state <- PatchState{Error: err}
			return
		}
		// Download HD manifest from patch repository, we'll use it multiple times.
		hdManifest, err := s.getManifest("hd/manifest.json")
		if err != nil {
			state <- PatchState{Error: err}
			return
		}

		for _, game := range conf.Games {
			// If the user has chosen to override the maphack config with their own,
			// we need to make sure the config is being ignored from the patch, and also
			// when reseting the maphack patch.
			var ignoredMaphackFiles []string

			if game.OverrideBHCfg {
				ignoredMaphackFiles = append(ignoredMaphackFiles, "BH.cfg")
			}

			// If maphack is disabled, make sure no rogue files have managed to stay in the directory.
			if !game.Maphack {
				installed, err := isMaphackInstalled(game.Location)
				if err != nil {
					state <- PatchState{Error: err}
					return
				}

				// If maphack is installed, but was supposed to be disabled, reset the patch.
				if installed {
					err := s.resetPatch(game.Location, maphackManifest.Files, ignoredMaphackFiles)
					if err != nil {
						state <- PatchState{Error: err}
						return
					}
				}
			}

			// If HD is disabled, make sure no rogue files have managed to stay in the directory.
			if !game.HD {
				installed, err := isHDInstalled(game.Location)
				if err != nil {
					state <- PatchState{Error: err}
					return
				}

				// If HD is installed, but was supposed to be disabled, reset the patch.
				if installed {
					err := s.resetPatch(game.Location, hdManifest.Files, nil)
					if err != nil {
						state <- PatchState{Error: err}
						return
					}
				}
			}

			// The install has been reset, let's validate the 1.13c version and apply missing files.
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
				err = s.applyMaphack(game.Location, state, progress, maphackManifest.Files, ignoredMaphackFiles)
				if err != nil {
					state <- PatchState{Error: err}
					return
				}
			}

			if game.HD {
				err = s.applyHDMod(game.Location, state, progress, hdManifest.Files)
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

// ApplyDEP will run  data execution prevention (DEP) on the Game.exe in the path.
func (s *Service) ApplyDEP(path string) error {
	// Run OS specific fix.
	return applyDEP(path)
}

func (s *Service) apply113c(path string, state chan PatchState, progress chan float32) error {
	state <- PatchState{Message: "Checking game version..."}

	// Download manifest from patch repository.
	manifest, err := s.getManifest("1.13c/manifest.json")
	if err != nil {
		return err
	}

	// Figure out which files to patch.
	patchFiles, patchLength, err := s.getFilesToPatch(manifest.Files, path, nil)
	if err != nil {
		return err
	}

	if len(patchFiles) > 0 {
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
	state <- PatchState{Message: "Checking Slashdiablo patch..."}

	// Download manifest from patch repository.
	manifest, err := s.getManifest("current/manifest.json")
	if err != nil {
		return err
	}

	// Figure out which files to patch.
	patchFiles, patchLength, err := s.getFilesToPatch(manifest.Files, path, nil)
	if err != nil {
		return err
	}

	if len(patchFiles) > 0 {
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

func (s *Service) applyMaphack(path string, state chan PatchState, progress chan float32, manifestFiles []PatchFile, ignoredFiles []string) error {
	state <- PatchState{Message: "Checking maphack..."}

	// Figure out which files to patch.
	patchFiles, patchLength, err := s.getFilesToPatch(manifestFiles, path, ignoredFiles)
	if err != nil {
		return err
	}

	if len(patchFiles) > 0 {
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

func (s *Service) applyHDMod(path string, state chan PatchState, progress chan float32, manifestFiles []PatchFile) error {
	// Update UI.
	state <- PatchState{Message: "Checking HD mod"}

	// Figure out which files to patch.
	patchFiles, patchLength, err := s.getFilesToPatch(manifestFiles, path, nil)
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

func (s *Service) getFilesToPatch(files []PatchFile, d2path string, filesToIgnore []string) ([]string, int64, error) {
	shouldPatch := make([]string, 0)
	var totalContentLength int64

	for _, file := range files {
		f := file

		// Check if the file should be ignored or not.
		if filesToIgnore != nil && len(filesToIgnore) > 0 {
			var ignore bool
			for _, ignored := range filesToIgnore {
				// If the current file should be ignored, just skip it.
				if f.Name == ignored {
					ignore = true
					break
				}
			}

			// File should be ignored, continue with the next.
			if ignore {
				continue
			}
		}

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
	s := &Service{
		githubClient:  githubClient,
		configService: configuration,
		logger:        logger,
		gameStates:    make(chan execState, 4),
	}

	// Setup game listener once, will stay alive for the duration
	// of the service's life cycle.
	go s.listenForGameStates()

	return s
}
