package model

import (
	"ImGollections/test"
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"testing"
)

func TestShouldCreateNewLayerFolderWith1Image(t *testing.T) {

	index := 5

	layerFolder, _ := newLayerFolder(test.BaseFolderImagePath, index)
	abs, _ := filepath.Abs(test.BaseFolderImagePath)
	assert.Equal(t, layerFolder.GetFolderPath(), abs)
	assert.Equal(t, layerFolder.GetLayerIndex(), index)
	assert.Len(t, layerFolder.GetImageArray(), 3)
	assert.Contains(t, layerFolder.GetImageArray()[0].GetImagePath(), filepath.Base(test.BaseImgPath))
}

func TestShouldCreateNewLayerFolderWith3Images(t *testing.T) {

	index := 5

	layerFolder, _ := newLayerFolder(test.Layer1FolderImagePath, index)
	abs, _ := filepath.Abs(test.Layer1FolderImagePath)
	assert.Equal(t, layerFolder.GetFolderPath(), abs)
	assert.Equal(t, layerFolder.GetLayerIndex(), index)
	assert.Len(t, layerFolder.GetImageArray(), 3)
	assert.Contains(t, layerFolder.GetImageArray()[0].GetImagePath(), filepath.Base(test.Layer11ImgPath))
	assert.Contains(t, layerFolder.GetImageArray()[1].GetImagePath(), filepath.Base(test.Layer12ImgPath))
	assert.Contains(t, layerFolder.GetImageArray()[2].GetImagePath(), filepath.Base(test.Layer13ImgPath))
}

func TestShouldCreateNewLayerFolderWith1ImageAnd1Skipped(t *testing.T) {

	index := 5

	layerFolder, _ := newLayerFolder(test.Layer2FolderImagePath, index)
	abs, _ := filepath.Abs(test.Layer2FolderImagePath)
	assert.Equal(t, layerFolder.GetFolderPath(), abs)
	assert.Equal(t, layerFolder.GetLayerIndex(), index)
	assert.Len(t, layerFolder.GetImageArray(), 3)
	assert.Contains(t, layerFolder.GetImageArray()[0].GetImagePath(), filepath.Base(test.Layer21ImgPath))
	assert.Contains(t, layerFolder.GetImageArray()[1].GetImagePath(), filepath.Base(test.Layer22ImgPath))
	assert.Contains(t, layerFolder.GetImageArray()[2].GetImagePath(), filepath.Base(test.Layer23ImgPath))
}

func TestShouldReturnErrorWithNegativeIndexWhileCreatingNewLayerFolder(t *testing.T) {

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	folderPath := "resources/hairs"
	index := -1

	layerFolder, err := newLayerFolder(folderPath, index)
	assert.Nil(t, layerFolder)
	assert.NotNil(t, err)
}

func TestShouldReturnTheExpectedListOfLayerFolder(t *testing.T) {

	baseLayerFolder, _ := newLayerFolder(test.BaseFolderImagePath, index)
	layer1LayerFolder, _ := newLayerFolder(test.Layer1FolderImagePath, index)
	layer2LayerFolder, _ := newLayerFolder(test.Layer2FolderImagePath, index)

	paths := []string{
		test.BaseFolderImagePath,
		test.Layer1FolderImagePath,
		test.Layer2FolderImagePath,
	}
	layerFolderSlice, err := CreateLayerFolders(paths, 0)

	assert.Nil(t, err)
	assert.Len(t, layerFolderSlice, 3)

	assert.Equal(t, layerFolderSlice[0].GetFolderPath(), baseLayerFolder.GetFolderPath())
	assert.Equal(t, 0, layerFolderSlice[0].GetLayerIndex())

	assert.Equal(t, layerFolderSlice[1].GetFolderPath(), layer1LayerFolder.GetFolderPath())
	assert.Equal(t, 1, layerFolderSlice[1].GetLayerIndex())

	assert.Equal(t, layerFolderSlice[2].GetFolderPath(), layer2LayerFolder.GetFolderPath())
	assert.Equal(t, 2, layerFolderSlice[2].GetLayerIndex())

}
