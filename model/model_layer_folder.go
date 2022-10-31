package model

import (
	"ImGollections/utils"
	"fmt"
	"github.com/pkg/errors"
	"os"
	"path/filepath"
	"strings"
)

type LayerFolder struct {
	folderPath      string
	layerIndex      int
	layerImageArray []LayerImage
}

func CreateLayerFolders(layerFolderPaths []string, layerIndex int) ([]LayerFolder, error) {

	if len(layerFolderPaths) == 0 {
		return nil, nil
	}

	layerFolder, err := newLayerFolder(layerFolderPaths[0], layerIndex)
	if err != nil {
		return nil, err
	}

	layerFolderSlice, err := CreateLayerFolders(layerFolderPaths[1:], layerIndex+1)
	if err != nil {
		return nil, err
	}

	return append([]LayerFolder{layerFolder}, layerFolderSlice...), nil
}

func newLayerFolder(folderPath string, layerIndex int) (LayerFolder, error) {

	if layerIndex < 0 {
		panic(errors.New("Received negative layer index during LayerFolder creation"))
	}

	isDir, err := utils.IsDir(folderPath)
	if err != nil {
		panic(err)
	}
	if !isDir {
		panic(errors.Wrap(err, folderPath+" is not a dir"))
	}

	fmt.Println("### Importing layer folder " + folderPath)

	children, _ := os.ReadDir(folderPath) // safe because checked before

	var layerImageArray []LayerImage
	for _, entry := range children {
		if !strings.HasSuffix(entry.Name(), ".png") {
			fmt.Println("*** Skipping image " + entry.Name() + ": it is not a png")
			continue
		}

		imagePath := filepath.Join(folderPath, entry.Name())
		image, err := NewImage(imagePath)
		if err != nil {
			fmt.Println("*** Skipping image " + imagePath + ": error during Image creation - " + err.Error())
			continue
		}

		layerImage, err := NewLayerImage(image)
		if err != nil {
			panic(errors.Wrap(err, "Error while creating LayerImage from image"))
		}
		layerImageArray = append(layerImageArray, layerImage)
	}

	return LayerFolder{
		folderPath,
		layerIndex,
		layerImageArray,
	}, nil
}

func (l LayerFolder) GetFolderPath() string {
	return l.folderPath
}

func (l LayerFolder) GetLayerIndex() int {
	return l.layerIndex
}

func (l LayerFolder) GetImageArray() []LayerImage {
	return l.layerImageArray
}
