package model

import (
	"ImGollections/test"
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

var index int
var baseLayerFolder LayerFolder

func TestMain(m *testing.M) {
	setup()
	exitCode := m.Run()
	os.Exit(exitCode)
}

func setup() {
	index = 5
	baseLayerFolder, _ = newLayerFolder(test.BaseFolderImagePath, index)
}

func TestShouldCreateNewLayerFolderListWithValidLayerFolder(t *testing.T) {

	layerFolderList, _ := newLayerFolderList(baseLayerFolder)

	assert.Len(t, layerFolderList.GetList(), 1)
	assert.Equal(t, layerFolderList.GetList()[0], baseLayerFolder)
}

func TestShouldCreateLayerFolderList(t *testing.T) {
	layerFolderList, err := CreateLayerFolderList(test.LayerFolderPaths)

	assert.Nil(t, err)
	assert.Len(t, layerFolderList.GetList(), 3)

	for i := range layerFolderList.GetList() {
		abs, _ := filepath.Abs(test.LayerFolderPaths[i])
		assert.Equal(t, abs, layerFolderList.GetList()[i].GetFolderPath())
		assert.Equal(t, i, layerFolderList.GetList()[i].GetLayerIndex())
	}
}
