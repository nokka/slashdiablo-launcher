package d2

import (
	"encoding/hex"
	"errors"
	"hash/crc32"
	"io"
	"os"
)

const (
	// Default polynomial when hashing.
	polynomial = 0xedb88320
)

var (
	// ErrCRCFileNotFound is used when the file to be hashed didn't exist.
	ErrCRCFileNotFound = errors.New("file not found")
)

//  hashCRC32 will load the file on the given file path, hash it and return sum as a string.
func hashCRC32(filePath string, polynomial uint32) (string, error) {
	// Open file that should be hashed.
	file, err := os.Open(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return "", ErrCRCFileNotFound
		}
		return "", err
	}

	// Close file when we're done.
	defer file.Close()

	// Create a table with the given polynomial to hash with.
	table := crc32.MakeTable(polynomial)

	// Create a new hasher.
	hash := crc32.New(table)

	// Copy the contents of the file into the hasher.
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	// Get the sum of the hash.
	hashInBytes := hash.Sum(nil)[:]

	// Encode the hash to a string.
	return hex.EncodeToString(hashInBytes), nil
}
