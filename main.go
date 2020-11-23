package main

import (
	img "image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"

	"github.com/nfnt/resize"
)

var maxIter int = 512
var xResolution uint = 2560
var yResolution uint = 1440

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

func getColor(nIterations int) (uint8, uint8, uint8) {
	var r, g, b uint8
	var lim1 int = 20
	var lim2 int = 40
	var floatMax float32 = 255.
	var floatNIter float32 = float32(nIterations)

	if nIterations <= 1 {
		r, g, b = 0, 0, 0
	} else if nIterations <= lim1 {
		g = uint8(floatMax * floatNIter / float32(lim1))
		r, b = 0, 0
	} else if nIterations <= lim2 {
		xv := uint8(floatMax * (floatNIter - 20) / float32(lim2-lim1))
		g = 255 - xv
		b = xv
		r = 0
	} else {
		xv := uint8(floatMax * (floatNIter - 40) / float32(maxIter-lim2))
		g = 0
		b = 255 - xv
		r = xv
	}
	return r, g, b
}

func main() {
	ptsPerPixel := 5
	width := int(xResolution) * ptsPerPixel
	height := int(yResolution) * ptsPerPixel

	image := img.NewRGBA(img.Rectangle{
		img.Point{0, 0},
		img.Point{width, height},
	})

	var xmin float64 = -2.6
	var xmax float64 = 1.5

	xDistPerPoint := (xmax - xmin) / float64(xResolution)
	var ymin float64 = -xDistPerPoint * float64(yResolution) / 2
	var ymax float64 = xDistPerPoint * float64(yResolution) / 2

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
			r, g, b := getColor(nIterations)
			image.Set(w, h, color.RGBA{r, g, b, 255})
		}
	}

	// Resizing

	imageResized := resize.Resize(xResolution, yResolution, image, resize.Lanczos3)

	outFile, _ := os.Create("testImage.png")
	png.Encode(outFile, imageResized)
}
