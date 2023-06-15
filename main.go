package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"os"
)

func main() {

	// Generate Random data
	data := GenerateData(400)
	CreateImage(810, 810, data)

}

type Vec2 struct {
	X, Y int
}

func GenerateData(dataCount int) []int {
	//var data []int
	data := make([]int, dataCount)
	for i := 1; i < dataCount; i++ {
		data[i] = rand.Intn(799)
		data[i]++
	}
	return data
}

var imgDrawn = 0

func CreateImage(width, height int, data []int) (bool, *image.RGBA) {

	// Create colors
	//red := color.RGBA{255, 50, 50, 0xFF}
	blue := color.RGBA{50, 180, 255, 0xFF}

	// Transform Data
	scale := Vec2{4, 10}
	barSpacing := 0

	var IsSorted = false
	var im *image.RGBA
	for !IsSorted {
		// Create image
		upLeft := image.Point{0, 0}
		lowRight := image.Point{width, height}
		img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

		pos := Vec2{10, 10}
		offset := Vec2{-5, -1}

		// For each rectangle
		for g := 1; g < len(data); g++ {

			scale.Y = data[g]

			// for each pixel
			for x := 0; x <= width; x++ {
				for y := 0; y < height; y++ {
					DrawRect(*img,
						x, y, // # top left
						pos.X+offset.X, pos.Y+offset.Y, // # bottom right + offset
						scale.X, scale.Y, // # scale of rectangle
						blue) // # color of rectangle

				}
			}
			// EOF each pixel
			offset.X += (scale.X / 2) + barSpacing

		} // EOF Image creation

		data, IsSorted = SortData(data)
		imgDrawn++

		// Image export
		MakeImage(img, fmt.Sprintf("%v", imgDrawn))
		im = img
	}
	return IsSorted, im

}

func MakeImage(img *image.RGBA, fileName string) {
	name := fmt.Sprintf("output/%s.png", fileName)
	f, _ := os.Create(name)
	png.Encode(f, img)
}

func DrawRect(img image.RGBA, x, y, posX, posY, scaleX, scaleY int, color color.RGBA) {
	if x > posX && x < (posX)+(scaleX/2) && y > posY && y < (posY)+(scaleY/2) {
		img.Set(x, y, color)
	}
}

func SortData(d []int) ([]int, bool) {

	var newData []int

	IsSorted := false
	swapCount := 0

	for !IsSorted {

		// For each number
		for i := 1; i < len(d); i++ {

			if d[i-1] > d[i] {
				// swap numbers
				tmp := d[i-1]
				d[i-1] = d[i]
				d[i] = tmp
				swapCount++
			}

		}

		if swapCount == 0 {
			// sorted
			newData = d
			fmt.Println("Data sorted -->>\n", newData)
			break
			//return newData, true
		}
		if swapCount > 0 {
			// not sorted
			newData = d
			return newData, false
		}

	}
	return newData, true
}

func GenerateGif(Image []*image.Paletted, Delay []int) {

	//Config := Image.Config

}
