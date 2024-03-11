//go:build windows

package fsutil

import (
	"debug/pe"
	"os"
)

func isExecutable(filepath string) bool {
	// Open the file
	file, err := os.Open(Abs(filepath))
	if err != nil {
		return false
	}
	defer file.Close()

	// Attempt to parse the PE file
	peFile, err := pe.NewFile(file)
	if err != nil {
		return false
	}

	// Verify the PE signature
	if peFile.Machine == pe.IMAGE_FILE_MACHINE_UNKNOWN {
		return false
	}

	// Check for "MZ" header
	mzHeader := make([]byte, 2)
	_, err = file.ReadAt(mzHeader, 0)
	if err != nil {
		return false
	}
	if string(mzHeader) != "MZ" {
		return false
	}

	// If all checks pass, return true
	return true
}
