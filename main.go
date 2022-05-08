package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	"os"
)

func main() {
	path := os.Args[1]
	fmt.Println(path, "PATH")
	infile, err := os.Open(path)


	if err != nil {
		panic(err)
	}
	defer infile.Close()


	img, _, err  := image.Decode(infile)

	x:= img.Bounds().Max.X
	y := img.Bounds().Max.Y
	
	
	fmt.Println(x, y)
}