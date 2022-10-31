package model

import (
	"ImGollections/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShouldCreateNewBaseLayer(t *testing.T) {

	ind := 1
	image, _ := NewImage(test.ZombieImgPath)
	layerImage, _ := NewLayerImage(image)
	layer := NewLayer(ind, layerImage)
	baseLayer := NewBaseLayer(layer)

	assert.Equal(t, baseLayer.GetLayerImage().GetImagePath(), test.ZombieImgPath)
	assert.Equal(t, baseLayer.GetLayerImage().GetRare(), false)
}

//func TestShouldReturnErrorWhileCreatingNewBaseLayerWithNilLayer(t *testing.T) {
//
//	baseLayer, err := NewBaseLayer(nil)
//	assert.NotNil(t, err)
//	assert.Nil(t, baseLayer)
//}
