package model

import (
	"ImGollections/test"
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"testing"
)

func TestShouldCreateNewBaseLayer(t *testing.T) {

	ind := 1
	abs, _ := filepath.Abs(test.BaseImgPath)
	image, _ := NewImage(abs)
	layerImage, _ := NewLayerImage(image)
	layer := NewLayer(ind, layerImage)
	baseLayer := NewBaseLayer(layer)

	assert.Equal(t, abs, baseLayer.GetLayerImage().GetImagePath())
	assert.Equal(t, false, baseLayer.GetLayerImage().GetRare())
}
