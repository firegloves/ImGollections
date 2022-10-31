package model

import (
	"ImGollections/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShouldCreateNewImageWithExistingImagePath(t *testing.T) {

	image, _ := NewImage(test.ZombieImgPath)
	assert.Equal(t, image.GetImagePath(), test.ZombieImgPath)
}

func TestNewImageShouldReturnErrorWithNonExistingImagePath(t *testing.T) {

	imagePath := "/Users/firegloves/workspace/NFTGonarator/resources/base/not_existing.png"
	_, err := NewImage(imagePath)
	assert.NotNil(t, err)
}

func TestNewImageShouldReturnErrorWithDirImagePath(t *testing.T) {

	_, err := NewImage(test.BaseFolderImagePath)
	assert.NotNil(t, err)
}
