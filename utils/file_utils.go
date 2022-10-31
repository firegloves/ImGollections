package utils

import (
	"errors"
	"os"
)

// FileExists returns whether the given file or directory exists
func FileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// IsDir returns whether the given path is a dir
func IsDir(path string) (bool, error) {
	if fileInfo, err := os.Stat(path); os.IsNotExist(err) {
		return false, errors.New("dir does not exist " + path)
	} else {
		return fileInfo.IsDir(), nil
	}
}

// IsFile returns whether the given path is a file
func IsFile(path string) (bool, error) {
	if fileInfo, err := os.Stat(path); os.IsNotExist(err) {
		return false, errors.New("dir does not exist" + path)
	} else {
		return !fileInfo.IsDir(), nil
	}
}
