// +build windows

package d2

import (
	"bufio"
	"bytes"
	"crypto/sha1"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
	"syscall"
	"unicode/utf8"

	"golang.org/x/sys/windows/registry"
)

// SHA1 of the different versions of Diablo Game.exe.
var hashList = map[string]string{
	"a875b98fa3a8b9300bcc04c84be1fa057eb277b5": "1.12",
	"af2b33c90b50ede8d9a8bca9b8d9720c87f78641": "1.13c",
	"11cd918cb6906295769d9be1b3e349e02af6b229": "1.13d",
	"3e64f12c6ef72847f49d301c2472280d4460589d": "1.14a",
	"11e940266c6838414c2114c2172227f982d4054e": "1.14b",
	"255691dd53e3bcd646e5c6e1e2e7b16da745b706": "1.14c",
	"af0ea93d2a652ceb11ac01ee2e4ae1ef613444c2": "1.14d",
}

// validate113cVersion will check the given installations Diablo II version.
func validate113cVersion(path string) (bool, error) {
	// Open local Game.exe.
	content, err := ioutil.ReadFile(localizePath(path) + "\\Game.exe")
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}

	// Hash the content of the Game.exe.
	hashed := fmt.Sprintf("%x", sha1.Sum(content))

	// Check the game version.
	version, ok := hashList[hashed]

	// Unknown game version.
	if !ok {
		return false, nil
	}

	return version == "1.13c", nil
}

// launch will execute the Diablo II.exe in the given directory.
func launch(path string, done chan execState) (*int, error) {
	// Localize the path.
	localized := localizePath(path)

	// Exec the Diablo II.exe.
	cmd := exec.Command(localized+"\\Diablo II.exe", "-w")
	cmd.Dir = localized

	// Collect the output from the command.
	var stderr bytes.Buffer

	// Pipe errors to our buffer.
	cmd.Stderr = &stderr

	fmt.Println("Starting...")
	if err := cmd.Start(); err != nil {
		return nil, err
	}

	fmt.Println("Started with pid:", cmd.Process.Pid)

	// Wait on separate thread.
	go func() {
		fmt.Println("Waiting...")

		if err := cmd.Wait(); err != nil {
			if exiterr, ok := err.(*exec.ExitError); ok {
				// The program has exited with an exit code != 0
				if status, ok := exiterr.Sys().(syscall.WaitStatus); ok {
					done <- execState{pid: &cmd.Process.Pid, err: fmt.Errorf("Exit status: %d : %s", status.ExitStatus(), stderr.String())}
				}
			} else {
				// Was some other wait error such as permissions, return the err.
				done <- execState{pid: &cmd.Process.Pid, err: fmt.Errorf("cmd.Wait: %d : %s", err, stderr.String())}
			}
		}

		fmt.Println("Waiting done...")
		done <- execState{pid: &cmd.Process.Pid, err: nil}
	}()

	return &cmd.Process.Pid, nil
}

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

// setGateway will set the gateway for Diablo II.
func setGateway(gateway string) error {
	var gatewayHex []byte
	var bnetIP string

	switch gateway {
	case GatewaySlashdiablo:
		bnetIP = SlashDiabloIP
		gatewayHex = getSlashGateway()
	case GatewayBattleNet:
		bnetIP = BattleNetIP
		gatewayHex = getBattleNetGateway()
	}

	// Open the Battle.net configuration registry directory.
	gatewayKey, err := registry.OpenKey(registry.CURRENT_USER, `Software\Battle.net\Configuration`, registry.QUERY_VALUE|registry.SET_VALUE)
	if err != nil {
		fmt.Println("OPEN KEY ERROR", err)
		return err
	}

	// Set the gateway hex.
	if err := gatewayKey.SetBinaryValue("Diablo II Battle.net Gateways", gatewayHex); err != nil {
		fmt.Println("SETTING ERROR", err)
		return err
	}

	// Close the registry when we're done.
	if err := gatewayKey.Close(); err != nil {
		return err
	}

	// Open the Battle net configuration directory.
	confKey, err := registry.OpenKey(registry.CURRENT_USER, `Software\Blizzard Entertainment\Diablo II`, registry.QUERY_VALUE|registry.SET_VALUE)
	if err != nil {
		fmt.Println("OPEN KEY ERROR", err)
		return err
	}

	// Set Battle.net IP.
	if err := confKey.SetStringValue("BNETIP", bnetIP); err != nil {
		fmt.Println("SETTING ERROR", err)
		return err
	}

	// Set the Command line args when starting.
	if err := confKey.SetStringValue("CmdLine", "-skiptobnet"); err != nil {
		fmt.Println("SETTING ERROR", err)
		return err
	}

	// Close the registry when we're done.
	if err := confKey.Close(); err != nil {
		return err
	}

	return nil
}

func isHDInstalled(path string) (bool, error) {
	filePath := localizePath(fmt.Sprintf("%s/%s", path, "D2HD.dll"))

	fmt.Println("HD FILE PATH", filePath)

	// Check if the file exists on disk.
	_, err := os.Stat(filePath)
	if err != nil {
		// File didn't exist on disk, return false.
		if os.IsNotExist(err) {
			return false, nil
		}
		// Unknown error.
		return false, err
	}

	return true, nil
}

func isMaphackInstalled(path string) (bool, error) {
	filePath := localizePath(fmt.Sprintf("%s/%s", path, "BH.dll"))

	fmt.Println("MAPHACK FILE PATH", filePath)

	// Check if the file exists on disk.
	_, err := os.Stat(filePath)
	if err != nil {
		// File didn't exist on disk, return false.
		if os.IsNotExist(err) {
			return false, nil
		}
		// Unknown error.
		return false, err
	}

	return true, nil
}

// localizePath will localize the path for the OS.
func localizePath(path string) string {
	// Windows uses backslashes for paths, so we'll reverse them.
	reversed := strings.Replace(path, "/", "\\", -1)

	// Remove the heading backslash.
	_, i := utf8.DecodeRuneInString(reversed)

	return reversed[i:]
}

func getSlashGateway() []byte {
	return []byte{
		0x31, 0x30, 0x30, 0x32,
		0x00, 0x30, 0x31, 0x00,
		0x70, 0x6c, 0x61, 0x79,
		0x2e, 0x73, 0x6c, 0x61,
		0x73, 0x68, 0x64, 0x69,
		0x61, 0x62, 0x6c, 0x6f,
		0x2e, 0x6e, 0x65, 0x74,
		0x00, 0x2d, 0x36, 0x00,
		0x73, 0x6c, 0x61, 0x73,
		0x68, 0x00, 0x65, 0x76,
		0x6e, 0x74, 0x2e, 0x73,
		0x6c, 0x61, 0x73, 0x68,
		0x64, 0x69, 0x61, 0x62,
		0x6c, 0x6f, 0x2e, 0x6e,
		0x65, 0x74, 0x00, 0x38,
		0x00, 0x45, 0x76, 0x65,
		0x6e, 0x74, 0x00, 0x00,
	}
}

func getBattleNetGateway() []byte {
	return []byte{
		0x31, 0x30, 0x30, 0x32,
		0x00, 0x30, 0x31, 0x00,
		0x75, 0x73, 0x77, 0x65,
		0x73, 0x74, 0x2e, 0x62,
		0x61, 0x74, 0x74, 0x6c,
		0x65, 0x2e, 0x6e, 0x65,
		0x74, 0x00, 0x38, 0x00,
		0x55, 0x2e, 0x53, 0x2e,
		0x20, 0x57, 0x65, 0x73,
		0x74, 0x00, 0x75, 0x73,
		0x65, 0x61, 0x73, 0x74,
		0x2e, 0x62, 0x61, 0x74,
		0x74, 0x6c, 0x65, 0x2e,
		0x6e, 0x65, 0x74, 0x00,
		0x36, 0x00, 0x55, 0x2e,
		0x53, 0x2e, 0x20, 0x45,
		0x61, 0x73, 0x74, 0x00,
		0x61, 0x73, 0x69, 0x61,
		0x2e, 0x62, 0x61, 0x74,
		0x74, 0x6c, 0x65, 0x2e,
		0x6e, 0x65, 0x74, 0x00,
		0x2d, 0x39, 0x00, 0x41,
		0x73, 0x69, 0x61, 0x00,
		0x65, 0x75, 0x72, 0x6f,
		0x70, 0x65, 0x2e, 0x62,
		0x61, 0x74, 0x74, 0x6c,
		0x65, 0x2e, 0x6e, 0x65,
		0x74, 0x00, 0x2d, 0x31,
		0x00, 0x45, 0x75, 0x72,
		0x6f, 0x70, 0x65, 0x00,
		0x00,
	}
}
