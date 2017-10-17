package linear

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
)

var w, h int = 1000, 1000

var img = image.NewRGBA(image.Rect(0, 0, w, h)) // Create a rect image with w-100 h-100
var col color.Color                             // set color variable
var white color.Color

// The main algorithm with bresenham
func line(x0, y0, x1, y1 int) {
	dx := x1 - x0
	dy := y1 - y0
	d := (float64)(dy / dx)
	deltaerr := math.Abs(d)
	bias := 0.0
	y := y0
	for x := x0; x < x1; x++ {
		img.Set(x, y, col)
		bias += deltaerr
		if bias >= 0.5 {
			y++
			bias += 1.0
		}
	}

}

func Bresenham() {
	col = color.RGBA{255, 0, 0, 255}
	white = color.RGBA{255, 255, 255, 255}
	// draw.Draw(img, img.Bounds(), &image.Uniform{white}, image.ZP, draw.Src)
	// line above could be used to fill white color to the whole pic, but it did not works

	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			img.Set(x, y, white)
		}
	}

	line(100, 200, 200, 300)
	f, err := os.OpenFile("./graphic/linear.png", os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	png.Encode(f, img)
}
