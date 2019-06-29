// +build darwin

package d2

import "fmt"

// checkVersion will check the given installations Diablo II version.
func checkVersion(dir string) string {
	return "1.12"
}

// Exec will execute the Diablo II.exe in the given directory.
func Exec(path string) error {
	fmt.Println(path)
	return nil
}
