package model

import (
	"ImGollections/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShouldCreateNewLayer(t *testing.T) {

	ind := 1
	image, _ := NewImage(test.BaseImgPath)
	layerImage, _ := NewLayerImage(image)
	layer := NewLayer(ind, layerImage)

	assert.Equal(t, layer.GetLayerImage().GetImagePath(), test.BaseImgPath)
	assert.Equal(t, layer.GetLayerImage().GetRare(), false)
	assert.Equal(t, layer.GetLayerIndex(), ind)
}
