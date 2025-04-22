package utils

import (
	"log"
	"os"
)

func DoesFileNotExist(filepath string) (bool, error) {
	_, err := os.Stat(filepath)
	return os.IsNotExist(err), err
}

func FatalErrCheck(err error) {
	if err != nil {
		log.Default().Fatal(err)
	}
}
