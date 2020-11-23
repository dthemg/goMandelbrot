package main

import (
	img "image"
	"image/color"
	"image/png"
	"math"
	"math/cmplx"
	"os"
)

var maxIter int = 512

var myColorScheme = []color.Color{
	color.RGBA{14, 55, 15, 255},
	color.RGBA{47, 97, 48, 255},
	color.RGBA{138, 171, 25, 255},
	color.RGBA{154, 187, 27, 255},
}

func iterate(c complex128) int {
	var cIter complex128 = 0
	for i := 0; i < maxIter; i++ {
		if cmplx.Abs(cIter) > 2 {
			return i
		}
		cIter = cIter*cIter + c
	}
	return maxIter
}

func calcColor(nIterations int, nColors int) int {
	maxLog := math.Log(float64(maxIter))
	colorIdx := (nColors - 1) * int(math.Log(float64(nIterations))/maxLog)
	return colorIdx
}

func main() {

	width := 10000
	height := 10000

	image := img.NewRGBA(img.Rectangle{
		img.Point{0, 0},
		img.Point{width, height},
	})

	var xmin float64 = -1.2
	var xmax float64 = 1.2
	var ymin float64 = -1.2
	var ymax float64 = 1.2

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
			colorIdx := calcColor(nIterations, 4)
			image.Set(w, h, myColorScheme[colorIdx])
		}
	}

	outFile, _ := os.Create("testImage.png")
	png.Encode(outFile, image)
}
