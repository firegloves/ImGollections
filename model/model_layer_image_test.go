package model

import (
	"ImGollections/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShouldCreateNewLayerImage(t *testing.T) {

	image, _ := NewImage(test.ZombieImgPath)
	layerImage, err := NewLayerImage(image)

	assert.Nil(t, err)
	assert.Equal(t, layerImage.GetImagePath(), test.ZombieImgPath)
	assert.Equal(t, layerImage.GetRare(), false)
}

func TestShouldCreateNewRareLayerImage(t *testing.T) {

	image, _ := NewImage(test.Hat2ImgPath)
	layerImage, err := NewLayerImage(image)

	assert.Nil(t, err)
	assert.Equal(t, layerImage.GetImagePath(), test.Hat2ImgPath)
	assert.Equal(t, layerImage.GetRare(), true)
}

func TestShouldReturnErrorWhileCreatingNewLayerImageWithNilImage(t *testing.T) {

	layerImage, err := NewLayerImage(nil)

	assert.NotNil(t, err)
	assert.Nil(t, layerImage)
}
