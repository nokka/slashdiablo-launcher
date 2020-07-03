// +build darwin

package d2

import (
	"fmt"
)

const (
	// ModMaphackIdentifier is the identifier we use to look for installs of maphack.
	ModMaphackIdentifier = "BH.dll"

	// ModHDIdentifier is the identifier we use to look for installs of hd mod.
	ModHDIdentifier = "D2HD.dll"
)

// validate113cVersion will check the given installations Diablo II version.
func validate113cVersion(dir string) (bool, error) {
	return true, nil
}

// launch will execute the Diablo II.exe in the given directory.
func launch(path string, flags []string, done chan execState) (*int, error) {
	pid := 1
	return &pid, nil
}

// localizePath will localize the path for the OS.
func localizePath(path string) string {
	return path
}

// configureForOS will set specific configurations, such as compatibility mode.
func configureForOS(path string) error {
	return nil
}

// applyDEP will run a fix to disable DEP.
func applyDEP(path string) error {
	return nil
}

func setDiabloRegistryKeys() error {
	return nil
}

func isModInstalled(path string, identifier string, manifest *Manifest) (bool, error) {
	filePath := localizePath(fmt.Sprintf("%s/%s", path, identifier))

	// Get the checksum from the file on disk.
	hashed, err := hashCRC32(filePath, polynomial)
	if err != nil {
		// The file doesn't exist on disk, so it's not installed.
		if err == ErrCRCFileNotFound {
			return false, nil
		}

		return false, err
	}

	var crc string
	// File exists on disk, find the CRC.
	for _, f := range manifest.Files {
		if f.Name == identifier {
			crc = f.CRC
			break
		}
	}

	if crc == hashed {
		return true, nil
	}

	return false, nil
}
