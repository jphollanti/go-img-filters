package main

import (
	"image/color"

	"gonum.org/v1/gonum/mat"
)

func spatialFilter(pixels *[][]color.Color, kernel *mat.Dense) {
	ppixel := *pixels

	rows, col := kernel.Dims()
	offset := float64(rows / 2)
	kernelLength := float64(col)

	newImage := make([][]color.Color, len(ppixel))
	for i := 0; i < len(newImage); i++ {
		newImage[i] = make([]color.Color, len(ppixel[0]))
	}
	copy(newImage, ppixel)

	for x := offset; x < float64(len(ppixel))-offset; x++ {
		for y := offset; y < float64(len(ppixel[0]))-offset; y++ {
			newPixel := color.RGBA{}

			for a := 0.0; a < kernelLength; a++ {
				for b := 0.0; b < kernelLength; b++ {
					xn := x + a - offset
					yn := y + a - offset
					r, g, bb, aa := ppixel[int(xn)][int(yn)].RGBA()
					newPixel.R += uint8(float64(uint8(r)) * (kernel.At(int(a), int(b))))
					newPixel.G += uint8(float64(uint8(g)) * (kernel.At(int(a), int(b))))
					newPixel.B += uint8(float64(uint8(bb)) * (kernel.At(int(a), int(b))))
					newPixel.A += uint8(float64(uint8(aa)) * (kernel.At(int(a), int(b))))
				}
			}

			newImage[int(x)][int(y)] = newPixel
		}
	}
	*pixels = newImage
}
