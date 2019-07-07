// +build windows

package d2

import (
	"crypto/sha1"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"unicode/utf8"
)

// This is a sha1 of the 1.13c Game.exe for Windows.
const gameHash = "hash"

// validate113cVersion will check the given installations Diablo II version.
func validate113cVersion(path string) (bool, error) {
	h := sha1.New()

	fmt.Println(localizePath(path) + "\\Game.exe")

	// Open local Game.exe to hash it.
	content, err := ioutil.ReadFile(localizePath(path) + "\\Game.exe"))
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return
	}

	h.Write(content)
	sum := h.Sum(nil)

	fmt.Println(sum)

	return false, nil
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
