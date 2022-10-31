package model

type BaseLayer struct {
	Layer
}

func NewBaseLayer(layer Layer) BaseLayer {
	return BaseLayer{
		layer,
	}
}
