package model

import (
	"github.com/pkg/errors"
)

type LayerFolderList struct {
	list []LayerFolder
}

func CreateLayerFolderList(layerFolderPaths []string) (LayerFolderList, error) {
	layerFolders, err := CreateLayerFolders(layerFolderPaths, 0)
	if err != nil {
		panic(errors.Wrap(err, "Error while creating LayerFolderList"))
	}
	if len(layerFolders) < 1 {
		panic(errors.New("No layers identified. At least 1 layer is required"))
	}

	return newLayerFolderList(layerFolders[0], layerFolders[1:]...)
}

func newLayerFolderList(layerFolder LayerFolder, optionalLayerFolders ...LayerFolder) (LayerFolderList, error) {

	return LayerFolderList{
		append([]LayerFolder{layerFolder}, optionalLayerFolders...),
	}, nil
}

func (l LayerFolderList) GetList() []LayerFolder {
	return l.list
}
