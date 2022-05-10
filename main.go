package main

import (
	"image/jpeg"
	_ "image/jpeg"
	"os"
)

func main() {
	path := os.Args[1]
	img := fromFile(path)

	f, _ := os.Create("gray.jpg")
	defer f.Close()
	pixelatedImg := pixelateImg(img, 5)
	grayScaleImg := toGrayScale(&pixelatedImg)
	rows := grayImgToAscii(grayScaleImg)

	finalString := ""
	for _, row := range rows {

		finalString += row + "\n"

	}
	os.WriteFile("output.txt", []byte(finalString), 0644)
	jpeg.Encode(f, &pixelatedImg, nil)

}
