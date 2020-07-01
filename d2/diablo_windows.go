// +build windows

package d2

import (
	"bytes"
	"crypto/sha1"
	"fmt"
	"io/ioutil"
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
	"27ddadbc457affed122564ae7a4bd2223181e15a": "1.13c", // Custom 1.13c build with HD icon.
	"11cd918cb6906295769d9be1b3e349e02af6b229": "1.13d",
	"3e64f12c6ef72847f49d301c2472280d4460589d": "1.14a",
	"11e940266c6838414c2114c2172227f982d4054e": "1.14b",
	"255691dd53e3bcd646e5c6e1e2e7b16da745b706": "1.14c",
	"af0ea93d2a652ceb11ac01ee2e4ae1ef613444c2": "1.14d",
}

const (
	// ModMaphackIdentifier is the identifier we use to look for installs of maphack.
	ModMaphackIdentifier = "BH.dll"

	// ModHDIdentifier is the identifier we use to look for installs of hd mod.
	ModHDIdentifier = "D2HD.dll"

	errRegistryKeyNotFound = "The system cannot find the file specified."

	// RegistryConfiguration is the configuration key for Diablo II.
	RegistryConfiguration = `Software\Battle.net\Configuration`

	// RegistryDiablo is where all Diablo II specific values are.
	RegistryDiablo = `Software\Blizzard Entertainment\Diablo II`

	// RegistryLayers is where all data about execution resides, like DEP.
	RegistryLayers = `Software\Microsoft\Windows NT\CurrentVersion\AppCompatFlags\Layers`

	// RegistryPermissions is the required permissions needed for operations on our keys.
	RegistryPermissions = registry.QUERY_VALUE | registry.SET_VALUE
)

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
func launch(path string, flags []string, done chan execState) (*int, error) {
	// Localize the path.
	localized := localizePath(path)

	// Exec the Diablo II.exe with the given command line args.
	cmd := exec.Command(localized+"\\Diablo II.exe", flags...)
	cmd.Dir = localized

	// Collect the output from the command.
	var stderr bytes.Buffer

	// Pipe errors to our buffer.
	cmd.Stderr = &stderr

	if err := cmd.Start(); err != nil {
		return nil, err
	}

	// Wait on separate thread.
	go func() {
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

		done <- execState{pid: &cmd.Process.Pid, err: nil}
	}()

	return &cmd.Process.Pid, nil
}

// configureForOS will set specific configurations, such as compatibility mode.
func configureForOS(path string) error {
	// The key name is the localized path for the Diablo II directory.
	keyName := fmt.Sprintf("%s\\%s", localizePath(path), "Game.exe")

	// Open the compatibility key directory.
	compatibilityKey, err := registry.OpenKey(registry.CURRENT_USER,
		RegistryLayers,
		RegistryPermissions,
	)
	if err != nil {
		return err
	}

	// Set Windows XP Service Pack 2 compatibility mode.
	if err := compatibilityKey.SetStringValue(keyName, "~ WINXPSP2"); err != nil {
		return err
	}

	// Close the registry when we're done.
	if err := compatibilityKey.Close(); err != nil {
		return err
	}

	return nil
}

// applyDEP will run a fix to disable DEP.
func applyDEP(path string) error {
	// The key name is the localized path for the Diablo II directory.
	keyName := fmt.Sprintf("%s\\%s", localizePath(path), "Diablo II.exe")

	// Open the dep key directory.
	depKey, err := registry.OpenKey(registry.CURRENT_USER,
		RegistryLayers,
		RegistryPermissions,
	)
	if err != nil {
		return err
	}

	// Set the value to disable DEP.
	if err := depKey.SetStringValue(keyName, "DisableNXShowUI"); err != nil {
		return err
	}

	// Close the registry when we're done.
	if err := depKey.Close(); err != nil {
		return err
	}

	return nil
}

// setDiabloRegistryKeys will remove the registry for BNETIP and set CmdLine options.
func setDiabloRegistryKeys() error {
	err := setGenericRegistry()
	if err != nil {
		return err
	}

	err = setGatewayRegistry()
	if err != nil {
		return err
	}

	return nil
}

func setGenericRegistry() error {
	var (
		d2Key registry.Key
		err   error
	)

	// Open the generic Diablo II key.
	d2Key, err = registry.OpenKey(registry.CURRENT_USER, RegistryDiablo, RegistryPermissions)
	if err != nil && err.Error() == errRegistryKeyNotFound {
		// The key doesn't exist, we have to create it.
		d2Key, _, err = registry.CreateKey(registry.CURRENT_USER, RegistryDiablo, RegistryPermissions)
		if err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	bnetIP, _, err := d2Key.GetStringValue("BNETIP")

	// If getting the string value errors and the error is something other than not found, return err.
	if err != nil && err.Error() != errRegistryKeyNotFound {
		return err
	}

	if bnetIP != "" {
		// The user has the BNETIP set, let's remove it, not to mess with other installs.
		err = d2Key.DeleteValue("BNETIP")
		if err != nil {
			return err
		}
	}

	// Close the registry when we're done.
	if err := d2Key.Close(); err != nil {
		return err
	}

	return nil
}

func setGatewayRegistry() error {
	var (
		gatewayKey registry.Key
		err        error
	)

	// Open the Battle.net configuration registry directory to set gateway.
	gatewayKey, err = registry.OpenKey(registry.CURRENT_USER, RegistryConfiguration, RegistryPermissions)
	if err != nil && err.Error() == errRegistryKeyNotFound {
		// The key doesn't exist, we have to create it.
		gatewayKey, _, err = registry.CreateKey(registry.CURRENT_USER, RegistryConfiguration, RegistryPermissions)
		if err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	// Hex representation of the Slashdiablo registry entry.
	gatewayHex := getSlashGatewayHex()

	// Set the gateway hex.
	if err := gatewayKey.SetBinaryValue("Diablo II Battle.net Gateways", gatewayHex); err != nil {
		return err
	}

	// Close the registry when we're done.
	if err := gatewayKey.Close(); err != nil {
		return err
	}

	return nil
}

func isModInstalled(path string, identifier string, manifest *Manifest) (bool, error) {
	filePath := localizePath(fmt.Sprintf("%s/%s", path, identifier))

	// Get the checksum from the file on disk.
	hashed, err := hashCRC32(filePath, polynomial)
	if err != nil {
		// The file doesn't exist on disk, so it's not installed.
		if err == ErrCRCFileNotFound {
			return false, nil
		}

		return false, err
	}

	var crc string
	// File exists on disk, find the CRC.
	for _, f := range manifest.Files {
		if f.Name == identifier {
			crc = f.CRC
			break
		}
	}

	if crc == hashed {
		return true, nil
	}

	return false, nil
}

// localizePath will localize the path for the OS.
func localizePath(path string) string {
	// Windows uses backslashes for paths, so we'll reverse them.
	reversed := strings.Replace(path, "/", "\\", -1)

	// Remove the heading backslash.
	_, i := utf8.DecodeRuneInString(reversed)

	return reversed[i:]
}

func getSlashGatewayHex() []byte {
	return []byte{
		0x31, 0x30, 0x30, 0x32,
		0x00, 0x30, 0x31, 0x00,
		0x70, 0x6c, 0x61, 0x79,
		0x2e, 0x73, 0x6c, 0x61,
		0x73, 0x68, 0x64, 0x69,
		0x61, 0x62, 0x6c, 0x6f,
		0x2e, 0x6e, 0x65, 0x74,
		0x00, 0x2d, 0x36, 0x00,
		0x53, 0x6c, 0x61, 0x73,
		0x68, 0x20, 0x44, 0x69,
		0x61, 0x62, 0x6c, 0x6f,
		0x00, 0x00,
	}
}
