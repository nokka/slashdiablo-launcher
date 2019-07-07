// +build darwin

package d2

import "fmt"

// validate113cVersion will check the given installations Diablo II version.
func validate113cVersion(dir string) (bool, error) {
	return false, nil
}

// Exec will execute the Diablo II.exe in the given directory.
func Exec(path string) error {
	fmt.Println(path)
	return nil
}

// localizePath will localize the path for the OS.
func localizePath(path string) string {
	return path
}
