// +build windows

package d2

import (
	"os/exec"
	"strings"
	"unicode/utf8"
)

// validate113cVersion will check the given installations Diablo II version.
func validate113cVersion(path string) bool {
	return false
}

// Exec will execute the Diablo II.exe in the given directory.
func Exec(path string) error {
	// Localize the path.
	localized := localizePath(path)

	// Exec the Diablo II.exe.
	cmd := exec.Command(localized+"\\Diablo II.exe", "-w")

	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

// localizePath will localize the path for the OS.
func localizePath(path string) string {
	// Windows uses backslashes for paths, so we'll reverse them.
	reversed := strings.Replace(path, "/", "\\", -1)

	// Remove the heading backslash.
	_, i := utf8.DecodeRuneInString(reversed)

	return reversed[i:]
}

/*func trimFirstRune(s string) string {
	_, i := utf8.DecodeRuneInString(s)
	return s[i:]
}*/
