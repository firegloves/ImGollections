package model

import (
	"ImGollections/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShouldCreateNewImageWithExistingImagePath(t *testing.T) {

	image, _ := NewImage(test.BaseImgPath)
	assert.Equal(t, image.GetImagePath(), test.BaseImgPath)
}

func TestNewImageShouldReturnErrorWithNonExistingImagePath(t *testing.T) {

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	imagePath := "resources/base/not_existing.png"
	_, err := NewImage(imagePath)
	assert.NotNil(t, err)
}

func TestNewImageShouldReturnErrorWithDirImagePath(t *testing.T) {

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	_, err := NewImage(test.BaseFolderImagePath)
	assert.NotNil(t, err)
}
