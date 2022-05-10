package main

import "image"

func imgBounds(img image.Image) (int, int) {
	return img.Bounds().Max.X, img.Bounds().Max.Y
}

func grayBounds(gray image.Gray) (int, int) {
	return gray.Bounds().Max.X, gray.Bounds().Max.Y

}

func incByValueAndByOne(iValue, iOne, value int) (int, int) {
	return iValue + value, iOne + 1
}
