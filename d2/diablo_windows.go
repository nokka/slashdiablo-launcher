// +build windows

package d2

import (
	"bufio"
	"crypto/sha1"
	"fmt"
	"io"
	"io/ioutil"
	"log"
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

// launch will execute the Diablo II.exe in the given directory.
func launch(path string) error {
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

// runDEPFix will run a fix to disable DEP.
/*func runDEPFix(path string) error {
	// Localize the path.
	//localized := localizePath(path)
	go func() {
		cmd := exec.Command("cmd.exe", "/C", "call", `DEP_fix.bat`)
		cmd.Dir = `D:\Testing\Diablo II`

		var out bytes.Buffer
		var stderr bytes.Buffer
		var stdin io.Reader
		cmd.Stdout = &out
		cmd.Stderr = &stderr
		cmd.Stdin = stdin

		err := cmd.Run()
		if err != nil {
			fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		}

		buf, err := ioutil.ReadAll(stdin)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(buf)
	}()

	return nil
}*/

/*func runDEPFix(path string) error {
	// Localize the path.
	//localized := localizePath(path)
	go func() {
		cmd := exec.Command("cmd.exe", "/C", "call", `DEP_fix.bat`)
		cmd.Dir = `D:\Testing\Diablo II`

		r, w, err := os.Pipe()
		if err != nil {
			return
		}

		cmd.Stdout = w
		cmd.Stderr = w

		go func() {
			scanner := bufio.NewScanner(r)
			for scanner.Scan() {
				line := scanner.Text()
				fmt.Println("READING LINE")
				fmt.Println(line)
			}
		}()

		err = cmd.Run()
		if err != nil {
			fmt.Println("ERROR", err)
			return
		}
	}()

	return nil
}*/

func runDEPFix(path string) error {
	go func() {
		cmd := exec.Command("cmd.exe", "/C", "call", `DEP_fix.bat`)
		cmd.Dir = localizePath(path)

		// Capture stdin for the command, so we can send data on it.
		stdin, err := cmd.StdinPipe()
		if err != nil {
			log.Fatal(err)
		}

		r, w, err := os.Pipe()
		if err != nil {
			return
		}

		cmd.Stdout = w
		cmd.Stderr = w

		err = cmd.Start()
		if err != nil {
			log.Fatal(err)
		}

		go func() {
			scanner := bufio.NewScanner(r)
			for scanner.Scan() {
				line := scanner.Text()
				fmt.Println(line)
			}
		}()

		fmt.Println("DONE READING")
		_, err = io.WriteString(stdin, "Yes")
		if err != nil {
			log.Fatal(err)
		}

		_, err = io.WriteString(stdin, "")
		if err != nil {
			log.Fatal(err)
		}

	}()

	return nil
}
