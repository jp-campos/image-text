package main

import (
	"fmt"
	"image"
	"image/color"
)

var asciiScale = []rune("$@B%8&WM#*oahkbdpqwmZO0QLCJUYXzcvunxrjft/\\|()1{}[]?-_+~<>i!lI;:,\"^`'. ")

func toGrayScale(img image.Image) image.Gray {
	gray := image.NewGray(img.Bounds())
	xMax, yMax := imgBounds(img)

	for x := 0; x < xMax; x++ {
		for y := 0; y < yMax; y++ {
			gray.Set(x, y, img.At(x, y))
		}
	}
	return *gray
}

func pixelateGrayImg(img image.Gray, pixelSize int) image.Gray {
	fmt.Println("Enters pixelate")
	xMax, yMax := grayBounds(img)

	newXBound := xMax / pixelSize
	newYBound := yMax / pixelSize

	pixelatedGray := image.NewGray(image.Rect(0, 0, newXBound, newYBound))
	fmt.Println("New bounds", newXBound, newYBound)
	for x, xNew := 0, 0; x < xMax; x, xNew = incByValueAndByOne(x, xNew, pixelSize) {
		for y, yNew := 0, 0; y < yMax; y, yNew = incByValueAndByOne(y, yNew, pixelSize) {
			fmt.Println("New img pos", xNew, yNew)
			yColor := avgPixelsGray(x, y, pixelSize, img)
			fmt.Println("Ycolor", yColor)
			pixelatedGray.Set(xNew, yNew, yColor)

		}
	}
	return *pixelatedGray
}

func pixelateImg(img image.Image, pixelSize int) image.RGBA {

	xMax, yMax := imgBounds(img)

	newXBound := xMax / pixelSize
	newYBound := yMax / pixelSize

	pixelated := image.NewRGBA(image.Rect(0, 0, newXBound, newYBound))

	for x, xNew := 0, 0; x < xMax; x, xNew = incByValueAndByOne(x, xNew, pixelSize) {
		for y, yNew := 0, 0; y < yMax; y, yNew = incByValueAndByOne(y, yNew, pixelSize) {

			var color color.Color

			if gray, ok := img.(*image.Gray); ok {
				color = avgPixelsGray(x, y, pixelSize, *gray)
			} else {
				color = avgPixelsColor(x, y, pixelSize, img)
			}
			//	fmt.Println(color)
			pixelated.Set(xNew, yNew, color)

		}
	}
	return *pixelated
}

func avgPixelsGray(xIni, yIni, pixelSize int, img image.Gray) color.Color {

	sum := 0
	for x, y := xIni, yIni; x < xIni+pixelSize; x++ {
		for ; y < yIni+pixelSize; y++ {
			sum += int(img.GrayAt(x, y).Y)
		}
	}

	var avg uint8 = uint8(sum / (pixelSize * pixelSize))

	return color.Gray{Y: avg}
}

func avgPixelsColor(xIni, yIni, blockSize int, img image.Image) color.Color {

	sumR := 0
	sumG := 0
	sumB := 0
	count := 0
	for x, y := xIni, yIni; x < xIni+blockSize; x++ {
		for ; y < yIni+blockSize; y++ {

			r, g, b, _ := img.At(x, y).RGBA()
			sumR, sumG, sumB = sumR+int(r/257), sumG+int(g/257), sumB+int(b/257)
			count++
		}
	}
	var avgR uint8 = uint8(sumR / count)
	var avgG uint8 = uint8(sumG / count)
	var avgB uint8 = uint8(sumB / count)

	return color.RGBA{avgR, avgG, avgB, uint8(255)}
}

func grayImgToAscii(img image.Gray) []string {
	asciiSize := len(asciiScale)
	xMax, yMax := grayBounds(img)

	asciiRows := make([]string, xMax)
	for y := 0; y < yMax; y++ {
		currRow := make([]rune, xMax)
		for x := 0; x < xMax; x++ {
			c := img.GrayAt(x, y)
			r := grayPixelToAscii(c, asciiSize)

			currRow[x] = r
		}
		asciiRows[y] = string(currRow)
	}
	return asciiRows
}

func grayPixelToAscii(color color.Gray, scaleSize int) rune {
	//fmt.Println("SCALE SIZe", scaleSize)

	scaleSizeFrom0 := scaleSize - 1
	brightness := color.Y

	normalizedIndex := int(brightness) * scaleSizeFrom0 / 255

	flippedIndex := scaleSizeFrom0 - normalizedIndex

	return asciiScale[flippedIndex]
}
