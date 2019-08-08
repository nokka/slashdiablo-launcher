// +build linux

package d2

import "fmt"

func isHDInstalled(path string) (bool, error) {
	return false, nil
}

func isMaphackInstalled(path string) (bool, error) {
	return false, nil
}

// validate113cVersion will check the given installations Diablo II version.
func validate113cVersion(dir string) (bool, error) {
	return false, nil
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

// runDEPFix will run a fix to disable DEP.
func runDEPFix(path string) error {
	return nil
}
