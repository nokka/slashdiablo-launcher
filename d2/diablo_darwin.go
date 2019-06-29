// +build darwin

package d2

import "fmt"

// CheckVersion will check the given installations Diablo II version.
func CheckVersion(dir string) string {
	return "1.13c"
}

// Exec will execute the Diablo II.exe in the given directory.
func Exec(path string) error {
	fmt.Println(path)
	return nil
}
