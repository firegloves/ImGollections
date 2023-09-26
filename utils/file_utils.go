package utils

import (
	"errors"
	"fmt"
	"os"
)

const resourceFolder = "resources"

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
		return false, errors.New("dir does not exist " + path)
	} else {
		return !fileInfo.IsDir(), nil
	}
}

// CreateDirIfNonExist if the received path does not exist, it create the relative folder. in any other case an error is returned
func CreateDirIfNonExist(path string) (bool, error) {
	isDir, err := IsDir(path)

	if err != nil {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			return false, err
		}
		isDir = true
	}

	if !isDir {
		return false, errors.New(fmt.Sprintf("Path %s does exist but it's not a folder", path))
	}

	return true, nil
}
