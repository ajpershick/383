package exer10

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
)

func DrawCircle(outerRadius, innerRadius int, outputFile string) {
	rectangle := image.Rect(0, 0, 200, 200)
	grid := image.NewRGBA(rectangle)

	for x := 0; x < 200; x++ {
		for y := 0; y < 200; y++ {
			if math.Pow(float64(x-100), 2)+math.Pow(float64(y-100), 2) <= math.Pow(float64(outerRadius), 2) &&
				math.Pow(float64(x-100), 2)+math.Pow(float64(y-100), 2) >= math.Pow(float64(innerRadius), 2) {

				grid.Set(x, y, color.NRGBA{
					R: 0,
					G: 0,
					B: 0,
					A: 255,
				})
			} else {
				grid.Set(x, y, color.NRGBA{
					R: 255,
					G: 255,
					B: 255,
					A: 255,
				})
			}
		}
	}

	f, err := os.Create(outputFile)
	if err != nil {
		panic(err)
	}

	if err := png.Encode(f, grid); err != nil {
		f.Close()
		panic(err)
	}

	if err := f.Close(); err != nil {
		panic(err)
	}

}
