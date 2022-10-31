package utils

import (
	"os"
)

func ApplyColorTransformation(imagePath string) {
	infile, err := os.Open(imagePath)
	if err != nil {
		// replace this with real error handling
		panic(err)
	}
	defer infile.Close()

	//src, _, err := image.Decode(infile)
	//if err != nil {
	//	// replace this with real error handling
	//	panic(err)
	//}

}
