//go:build darwin

package fsutil

import "os"

func isExecutable(filepath string) bool {
	info, err := os.Stat(filepath)
	if err != nil {
		return false
	}

	return (info.Mode()&0111 != 0)
}
