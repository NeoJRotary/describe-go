package describe

import "os"

// FileExist is file exist ?
func FileExist(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}
