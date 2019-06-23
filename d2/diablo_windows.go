// +build windows

package d2

import (
	"os/exec"
)

// CheckVersion will check the given installations Diablo II version.
func CheckVersion(path string) string {
	return "1.13c"
}

// Exec will execute the Diablo II.exe in the given directory.
func Exec(path string) error {
	cmd := exec.Command(path, "-w")

	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
