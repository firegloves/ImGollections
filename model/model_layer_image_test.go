package model

import (
	"ImGollections/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShouldCreateNewLayerImage(t *testing.T) {

	image, _ := NewImage(test.BaseImgPath)
	layerImage, err := NewLayerImage(image)

	assert.Nil(t, err)
	assert.Equal(t, layerImage.GetImagePath(), test.BaseImgPath)
	assert.Equal(t, layerImage.GetRare(), false)
}

func TestShouldCreateNewLayer12Image(t *testing.T) {

	image, _ := NewImage(test.Layer12ImgPath)
	layerImage, err := NewLayerImage(image)

	assert.Nil(t, err)
	assert.Equal(t, layerImage.GetImagePath(), test.Layer12ImgPath)
}
