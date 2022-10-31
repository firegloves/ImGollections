package model

import (
	"ImGollections/utils"
	"github.com/pkg/errors"
)

type Image struct {
	imagePath string
}

func NewImage(imagePath string) (Image, error) {

	isFile, err := utils.IsFile(imagePath)
	if err != nil {
		panic(errors.Wrap(err, imagePath))
	}
	if !isFile {
		panic(errors.New(imagePath + " is not a file"))
	}

	return Image{
		imagePath,
	}, nil
}

func (i Image) GetImagePath() string {
	return i.imagePath
}
