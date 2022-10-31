package generator

import (
	"ImGollections/config"
	"ImGollections/model"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
)

func Generate(folderPaths []string) error {

	layerFolderList, err := model.CreateLayerFolderList(folderPaths)
	if err != nil {
		return err
	}

	layerImageMatrix := generateLayerImageMatrix(layerFolderList.GetList())
	layerMatrix := generateLayerMatrix(layerImageMatrix)
	fmt.Printf("COUNT %d\n", len(layerMatrix))
	pictureList := generateCollection(layerMatrix)
	saveCollection(pictureList)

	return nil
}

func generateLayerImageMatrix(layerFolder []model.LayerFolder) [][]model.LayerImage {

	var layerImageSlice [][]model.LayerImage

	maxInd := len(layerFolder) - 1

	// init matrix
	if maxInd == 0 {
		n := len(layerFolder[maxInd].GetImageArray())
		matrix := make([][]model.LayerImage, n)
		for i := 0; i < n; i++ {
			matrix[i] = []model.LayerImage{layerFolder[maxInd].GetImageArray()[i]}
		}
		return matrix
	} else {
		// call recursively
		innerLayerImageSlice := generateLayerImageMatrix(layerFolder[:len(layerFolder)-1])
		// for each layer slice (list of layers composing an image) obtained by the previous iteration
		for c := range innerLayerImageSlice {
			imageArray := layerFolder[len(layerFolder)-1].GetImageArray()
			// for each image present in the current layer folder
			for _, img := range imageArray {
				// append the current image to the list of the layer slices available
				newSlice := append([]model.LayerImage(nil), innerLayerImageSlice[c]...)
				images := append(newSlice, img)
				// append the new image slice to the total image slice available
				layerImageSlice = append(layerImageSlice, images)
			}
		}
		return layerImageSlice
	}
}

func generateLayerMatrix(matrix [][]model.LayerImage) [][]model.Layer {

	var layerMatrix [][]model.Layer

	for i := range matrix {
		layerImageList := matrix[i]
		if len(layerImageList) <= 1 {
			fmt.Println("*** Skipping layerImagelist n ", i)
			continue
		}

		var layerList []model.Layer
		for c := range layerImageList {
			layerList = append(layerList, model.NewLayer(i, layerImageList[c]))
		}
		layerMatrix = append(layerMatrix, layerList)
	}

	return layerMatrix
}

func generateCollection(layerMatrix [][]model.Layer) []model.Picture {

	var pictureList []model.Picture

	for i := range layerMatrix {
		pictureList = append(pictureList, model.NewPicture(model.NewBaseLayer(layerMatrix[i][0]), layerMatrix[i][1:]))
	}

	return pictureList
}

func saveCollection(pictureList []model.Picture) error {

	offset := 100

	desiredModule := len(pictureList) / config.GetConfig().Desiredpicturesnumber

	// https://medium.com/@arrafiv/basic-image-processing-with-go-combining-images-and-texts-8510d9214e55
	for i, pic := range pictureList {

		//fmt.Printf("ITER: %d - DESIRED %d - REST: %d\n", i, config.GetConfig().Desiredpicturesnumber, i%desiredModule)
		if i%desiredModule == 0 {

			fmt.Printf("########## Generating zombie %d\n", i)

			//create image’s background
			bgImg := image.NewRGBA(image.Rect(0, 0, 100, 100))
			//set the background color
			draw.Draw(bgImg, bgImg.Bounds(), &image.Uniform{C: color.Transparent}, image.Point{}, draw.Src)

			//fmt.Printf("**** Adding base layer %s\n", pic.GetBaseLayer().Layer.GetLayerImage().GetImagePath())
			// add base layer
			addLayerToImage(bgImg, pic.GetBaseLayer().Layer)
			// add remaining layers
			for _, layer := range pic.GetLayerList() {
				//fmt.Printf("**** Adding layer %s\n", layer.GetLayerImage().GetImagePath())
				addLayerToImage(bgImg, layer)
			}

			// save img
			savePNG(bgImg, fmt.Sprintf("/Users/firegloves/workspace/NFTGonarator/output/zombie_%d.png", i+offset))
		}
	}

	return nil
}

func addLayerToImage(img *image.RGBA, layer model.Layer) {
	// create layer
	layerFile, err := os.Open(layer.GetLayerImage().GetImagePath())
	if err != nil {
		panic("READ LAYER IMAGE ERROR")
	}
	imgLayer, _, err := image.Decode(layerFile)
	defer layerFile.Close()

	draw.Draw(img, img.Bounds(), imgLayer, image.Point{}, draw.Over)
}

func savePNG(img image.Image, filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	if err := png.Encode(f, img); err != nil {
		return err
	}
	if err := f.Close(); err != nil {
		return err
	}
	return nil
}
