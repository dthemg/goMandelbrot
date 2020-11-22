package main

import (
	img "image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func iterate(c complex128) int {
	var cIter complex128 = 0
	maxIter := 100
	for i := 0; i < maxIter; i++ {
		if cmplx.Abs(cIter) > 2 {
			return i
		}
		cIter = cIter*cIter + c
	}
	return maxIter
}

func main() {

	width := 1000
	height := 1000

	image := img.NewRGBA(img.Rectangle{
		img.Point{0, 0},
		img.Point{width, height},
	})

	var myColorScheme = []color.Color{
		color.RGBA{14, 55, 15, 255},
		color.RGBA{47, 97, 48, 255},
		color.RGBA{138, 171, 25, 255},
		color.RGBA{154, 187, 27, 255},
	}

	var xmin float64 = -5
	var xmax float64 = 5
	var ymin float64 = -5
	var ymax float64 = 5

	var input complex128
	var nIterations int
	widthFloat := float64(width)
	heightFloat := float64(height)
	for w := 0; w < width; w++ {
		for h := 0; h < height; h++ {
			wFloat := float64(w)
			hFloat := float64(h)
			xv := xmin + (xmax-xmin)*wFloat/widthFloat
			yv := ymin + (ymax-ymin)*hFloat/heightFloat
			nIterations = iterate(complex(xv, yv))

		}
	}

	outFile, _ := os.Create("testImage.png")
	png.Encode(outFile, image)
}
