package model

import (
	"ImGollections/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShouldCreateNewLayerFolderWith1Image(t *testing.T) {

	index := 5

	layerFolder, _ := newLayerFolder(test.BaseFolderImagePath, index)
	assert.Equal(t, layerFolder.GetFolderPath(), test.BaseFolderImagePath)
	assert.Equal(t, layerFolder.GetLayerIndex(), index)
	assert.Len(t, layerFolder.GetImageArray(), 1)
	assert.Contains(t, layerFolder.GetImageArray()[0].GetImagePath(), test.ZombieImgPath)
}

func TestShouldCreateNewLayerFolderWith2Images(t *testing.T) {

	index := 5

	layerFolder, _ := newLayerFolder(test.HatsFolderImagePath, index)
	assert.Equal(t, layerFolder.GetFolderPath(), test.HatsFolderImagePath)
	assert.Equal(t, layerFolder.GetLayerIndex(), index)
	assert.Len(t, layerFolder.GetImageArray(), 2)
	assert.Contains(t, layerFolder.GetImageArray()[0].GetImagePath(), test.Hat1ImgPath)
	assert.Contains(t, layerFolder.GetImageArray()[1].GetImagePath(), test.Hat2ImgPath)
}

func TestShouldCreateNewLayerFolderWith1ImageAnd1Skipped(t *testing.T) {

	index := 5

	layerFolder, _ := newLayerFolder(test.HairsFolderImagePath, index)
	assert.Equal(t, layerFolder.GetFolderPath(), test.HairsFolderImagePath)
	assert.Equal(t, layerFolder.GetLayerIndex(), index)
	assert.Len(t, layerFolder.GetImageArray(), 1)
	assert.Contains(t, layerFolder.GetImageArray()[0].GetImagePath(), test.Hair1ImgPath)
}

func TestShouldReturnErrorWithNegativeIndexWhileCreatingNewLayerFolder(t *testing.T) {

	folderPath := "/Users/firegloves/workspace/NFTGonarator/resources/hairs"
	index := -1

	layerFolder, err := newLayerFolder(folderPath, index)
	assert.Nil(t, layerFolder)
	assert.NotNil(t, err)
}

func TestShouldReturnTheExpectedListOfLayerFolder(t *testing.T) {

	paths := []string{
		test.BaseFolderImagePath,
		test.HairsFolderImagePath,
		test.HatsFolderImagePath,
	}
	layerFolderSlice, err := CreateLayerFolders(paths, 0)

	assert.Nil(t, err)
	assert.Len(t, layerFolderSlice, 3)

	for i := range layerFolderSlice {
		assert.Equal(t, layerFolder.GetFolderPath(), layerFolderSlice[i])
		assert.Equal(t, layerFolder.GetLayerIndex(), i)

		// TODO add images assert
	}
}
