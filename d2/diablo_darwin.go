// +build darwin

package d2

import "fmt"

// validate113cVersion will check the given installations Diablo II version.
func validate113cVersion(dir string) bool {
	return false
}

// Exec will execute the Diablo II.exe in the given directory.
func Exec(path string) error {
	fmt.Println(path)
	return nil
}
