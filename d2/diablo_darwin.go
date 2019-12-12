// +build darwin

package d2

import (
	"fmt"
	"os"
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

func isHDInstalled(path string, manifest *Manifest) (bool, error) {
	filePath := localizePath(fmt.Sprintf("%s/%s", path, "D2HD.dll"))

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
		if f.Name == "D2HD.dll" {
			crc = f.CRC
			break
		}
	}

	if crc == hashed {
		return true, nil
	}

	return false, nil
}

func isMaphackInstalled(path string) (bool, error) {
	filePath := localizePath(fmt.Sprintf("%s/%s", path, "BH.dll"))

	// Check if the file exists on disk.
	_, err := os.Stat(filePath)
	if err != nil {
		// File didn't exist on disk, return false.
		if os.IsNotExist(err) {
			return false, nil
		}
		// Unknown error.
		return false, err
	}

	return true, nil
}

// setGateway will set the gateway for Diablo II.
func setGateway(gateway string) error {
	return nil
}
