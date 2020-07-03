// +build linux

package d2

const (
	// ModMaphackIdentifier is the identifier we use to look for installs of maphack.
	ModMaphackIdentifier = "BH.dll"

	// ModHDIdentifier is the identifier we use to look for installs of hd mod.
	ModHDIdentifier = "D2HD.dll"
)

// validate113cVersion will check the given installations Diablo II version.
func validate113cVersion(dir string) (bool, error) {
	return false, nil
}

// launch will execute the Diablo II.exe in the given directory.
func launch(path string, flags []string, done chan execState) (*int, error) {
	id := 1
	return &id, nil
}

// localizePath will localize the path for the OS.
func localizePath(path string) string {
	return path
}

// configureForOS will set specific configurations, such as compatibility mode.
func configureForOS(path string) error {
	return nil
}

// applyDEP will run a fix to disable DEP.
func applyDEP(path string) error {
	return nil
}

func setDiabloRegistryKeys() error {
	return nil
}

func isModInstalled(path string, identifier string, manifest *Manifest) (bool, error) {
	return false, nil
}
