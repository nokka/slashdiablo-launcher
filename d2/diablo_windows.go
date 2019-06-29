// +build windows

package d2

import (
	"os/exec"
	"strings"
	"unicode/utf8"
)

// checkVersion will check the given installations Diablo II version.
func checkVersion(path string) string {
	return "1.12"
}

// Exec will execute the Diablo II.exe in the given directory.
func Exec(path string) error {
	// Windows uses backslashes for paths, so we'll reverse them.
	reversed := strings.Replace(path, "/", "\\", -1)

	// Remove the heading backslash on Windows.
	trimmed := trimFirstRune(reversed)

	// Exec the Diablo II.exe.
	cmd := exec.Command(trimmed+"\\Diablo II.exe", "-w")

	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func trimFirstRune(s string) string {
	_, i := utf8.DecodeRuneInString(s)
	return s[i:]
}
