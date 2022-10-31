package model

import (
	"ImGollections/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShouldCreateNewPictureWithOnlyBaseLayer(t *testing.T) {

	ind := 1
	image, _ := NewImage(test.ZombieImgPath)
	layerImage, _ := NewLayerImage(image)
	layer := NewLayer(ind, layerImage)
	baseLayer := NewBaseLayer(layer)
	picture := NewPicture(baseLayer, nil)

	assert.Equal(t, picture.GetBaseLayer(), baseLayer)
	assert.Nil(t, picture.GetLayerList())
}

func TestShouldCreateNewPictureWithMultipleLayers(t *testing.T) {

	ind := 1
	image, _ := NewImage(test.ZombieImgPath)
	layerImage, _ := NewLayerImage(image)
	layer := NewLayer(ind, layerImage)
	layerTwo := NewLayer(2, layerImage)
	baseLayer := NewBaseLayer(layer)
	picture := NewPicture(baseLayer, []*Layer{layerTwo})

	assert.Equal(t, picture.GetBaseLayer(), baseLayer)
	assert.Len(t, picture.GetLayerList(), 1)
	assert.Equal(t, picture.GetLayerList()[0], layerTwo)
}

//func TestShouldReturnErrorWhileCreatingNewPictureWithNullBaseLayer(t *testing.T) {
//
//	picture, err := NewPicture(nil, nil)
//
//	assert.NotNil(t, err)
//	assert.Nil(t, picture)
//}
