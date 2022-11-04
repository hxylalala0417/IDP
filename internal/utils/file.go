package utils

import "os"

// Exists checks whether the given path exists
func Exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}
