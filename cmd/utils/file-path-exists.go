package utils

import "os"

func FilePathExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}
