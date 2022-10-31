package generator

import (
	"ImGollections/model"
	"ImGollections/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShouldGenerateCollection(t *testing.T) {

	layerFolderList, _ := model.CreateLayerFolderList(test.LayerFolderPaths)
	collection := generateLayerImageMatrix(layerFolderList.GetList())

	assert.Equal(t, len(collection), 10)
}
