package model

type Layer struct {
	layerIndex int
	layerImage LayerImage
}

func CreateLayer(layerIndex int, layerImage LayerImage, colorVariation string) Layer {

	// TODO use color

	return NewLayer(layerIndex, layerImage)
}

func NewLayer(layerIndex int, layerImage LayerImage) Layer {

	if layerIndex < 0 {
		panic("Received negative layer index during Layer creation")
	}

	return Layer{
		layerIndex,
		layerImage,
	}
}

func (l Layer) GetLayerIndex() int {
	return l.layerIndex
}

func (l Layer) GetLayerImage() LayerImage {
	return l.layerImage
}
