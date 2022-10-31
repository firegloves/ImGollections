package model

import (
	"os"
	"strings"
)

type LayerImage struct {
	Image
	rare bool
}

func NewLayerImage(image Image) (LayerImage, error) {

	fileInfo, _ := os.Stat(image.GetImagePath())

	return LayerImage{
		Image{image.imagePath},
		strings.HasPrefix(fileInfo.Name(), "rare_"),
	}, nil
}

func (l LayerImage) GetRare() bool {
	return l.rare
}
