package utils

import "os"

func DoesFileNotExist(filepath string) bool {
	_, err := os.Stat(filepath)
	return os.IsNotExist(err)
}
