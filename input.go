package main

import (
	"image"
	"os"
)

func fromFile(path string) image.Image {

	infile, err := os.Open(path)

	if err != nil {
		panic(err)
	}
	defer infile.Close()

	img, _, err := image.Decode(infile)
	return img
}
