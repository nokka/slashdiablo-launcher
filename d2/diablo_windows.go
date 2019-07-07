// +build windows

package d2

import (
	"crypto/sha1"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"unicode/utf8"
)

// SHA1 of the 1.13c Game.exe for Windows.
const gameHash = "af2b33c90b50ede8d9a8bca9b8d9720c87f78641"

// validate113cVersion will check the given installations Diablo II version.
func validate113cVersion(path string) (bool, error) {
	// Open local Game.exe to hash it.
	content, err := ioutil.ReadFile(localizePath(path) + "\\Game.exe")
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}

	return fmt.Sprintf("%x", sha1.Sum(content)) == gameHash, nil
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
