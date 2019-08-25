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
func launch(path string, done chan execState) (*int, error) {
	fmt.Println(path)
	id := 1
	return &id, nil
}

// localizePath will localize the path for the OS.
func localizePath(path string) string {
	return path
}

// applyDEP will run a fix to disable DEP.
func applyDEP(path string) error {
	return nil
}

func isHDInstalled(path string) (bool, error) {
	filePath := localizePath(fmt.Sprintf("%s/%s", path, "D2HD.dll"))

	fmt.Println("HD FILE PATH", filePath)

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

func isMaphackInstalled(path string) (bool, error) {
	filePath := localizePath(fmt.Sprintf("%s/%s", path, "BH.dll"))

	fmt.Println("MAPHACK FILE PATH", filePath)

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
