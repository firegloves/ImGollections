package model

import (
	"ImGollections/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShouldCreateNewLayer(t *testing.T) {

	ind := 1
	image, _ := NewImage(test.ZombieImgPath)
	layerImage, _ := NewLayerImage(image)
	layer := NewLayer(ind, layerImage)

	assert.Equal(t, layer.GetLayerImage().GetImagePath(), test.ZombieImgPath)
	assert.Equal(t, layer.GetLayerImage().GetRare(), false)
	assert.Equal(t, layer.GetLayerIndex(), ind)
}

func TestShouldReturnErrorWhileCreatingNewLayerWithNilLayerImage(t *testing.T) {

	// TODO fix test

	//layer := NewLayer(1, nil)
	//assert.NotNil(t, err)
	//assert.Nil(t, layer)
}
