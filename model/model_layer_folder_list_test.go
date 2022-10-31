package model

import (
	"ImGollections/test"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var index int
var layerFolder *LayerFolder

func TestMain(m *testing.M) {
	setup()
	exitCode := m.Run()
	os.Exit(exitCode)
}

func setup() {
	index = 5
	layerFolder, _ = newLayerFolder(test.BaseFolderImagePath, index)
}

func TestShouldCreateNewLayerFolderListWithValidLayerFolder(t *testing.T) {

	layerFolderList, _ := newLayerFolderList(layerFolder)

	assert.Len(t, layerFolderList.GetList(), 1)
	assert.Equal(t, layerFolderList.GetList()[0], layerFolder)
}

func TestShouldReturnErrorWhenCreatingNewLayerFolderListWithNilLayerFolder(t *testing.T) {

	layerFolderList, err := newLayerFolderList(nil)

	assert.NotNil(t, err)
	assert.Nil(t, layerFolderList)
}

func TestShouldCreateLayerFolderList(t *testing.T) {
	layerFolderList, err := CreateLayerFolderList(test.LayerFolderPaths)

	assert.Nil(t, err)
	assert.Len(t, layerFolderList.GetList(), 3)

	for i := range layerFolderList.GetList() {
		assert.Equal(t, layerFolderList.GetList()[i].GetFolderPath(), test.LayerFolderPaths[i])
		assert.Equal(t, layerFolderList.GetList()[i].GetLayerIndex(), i)

		// TODO add images assert
	}
}

//func TestShouldAddAFolderLayerToAnExistingLayerFolderList(t *testing.T) {
//
//	layerFolderList, err := NewLayerFolderList(layerFolder)
//	layerFolder2, _ := newLayerFolder(test.HatsFolderImagePath, 3)
//	layerFolderListTwo, err := AddLayerFolder(layerFolderList, layerFolder2)
//
//	assert.Nil(t, err)
//	assert.NotNil(t, layerFolderListTwo)
//	assert.Len(t, layerFolderListTwo.GetList(), 2)
//	assert.Equal(t, layerFolderListTwo.GetList()[0], layerFolder)
//	assert.Equal(t, layerFolderListTwo.GetList()[1], layerFolder2)
//}
//
//func TestShouldAddTwoFolderLayersToAnExistingLayerFolderList(t *testing.T) {
//
//	layerFolderList, err := NewLayerFolderList(layerFolder)
//
//	layerFolder2, _ := newLayerFolder(test.HatsFolderImagePath, 5)
//	layerFolder3, _ := newLayerFolder(test.HairsFolderImagePath, 4)
//
//	layerFolderList, err = AddLayerFolder(layerFolderList, layerFolder2)
//	layerFolderList, err = AddLayerFolder(layerFolderList, layerFolder3)
//
//	assert.Nil(t, err)
//	assert.NotNil(t, layerFolderList)
//	assert.Len(t, layerFolderList.GetList(), 3)
//	assert.Equal(t, layerFolderList.GetList()[0], layerFolder)
//	assert.Equal(t, layerFolderList.GetList()[1], layerFolder2)
//	assert.Equal(t, layerFolderList.GetList()[2], layerFolder3)
//}
//
//func TestShouldAddTwoFolderLayersToAnExistingLayerFolderListNil(t *testing.T) {
//
//	layerFolderList, err := NewLayerFolderList(layerFolder)
//
//	layerFolderList, err = AddLayerFolder(layerFolderList, nil)
//
//	assert.NotNil(t, err)
//	assert.NotNil(t, layerFolderList)
//	assert.Len(t, layerFolderList.GetList(), 1)
//	assert.Equal(t, layerFolderList.GetList()[0], layerFolder)
//}
