package model

type Picture struct {
	baseLayer BaseLayer
	layerList []Layer
}

func NewPicture(baseLayer BaseLayer, layerList []Layer) Picture {

	return Picture{
		baseLayer,
		layerList,
	}
}

func (p Picture) GetBaseLayer() BaseLayer {
	return p.baseLayer
}

func (p Picture) GetLayerList() []Layer {
	return p.layerList
}
