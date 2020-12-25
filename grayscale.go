package main

import (
	"fmt"
	"image/color"
	"math"
	"sync"
)

type GrayScaleMode string

const (
	Luminosity GrayScaleMode = "luminosity"
	Average    GrayScaleMode = "average"
	Lightness  GrayScaleMode = "lightness"
)

func grayScale(pixels *Tensor, mode GrayScaleMode) {
	ppixels := *pixels
	xLen := len(ppixels)
	yLen := len(ppixels[0])
	//create new image
	newImage := make(Tensor, xLen)
	for i := 0; i < len(newImage); i++ {
		newImage[i] = make([]color.Color, yLen)
	}
	//idea is processing pixels in parallel
	wg := sync.WaitGroup{}
	for x := 0; x < xLen; x++ {
		for y := 0; y < yLen; y++ {
			wg.Add(1)
			go func(x, y int) {
				pixel := ppixels[x][y]
				originalColor, ok := color.RGBAModel.Convert(pixel).(color.RGBA)
				if !ok {
					fmt.Println("type conversion went wrong")
				}
				var grey uint8
				if mode == Luminosity {
					grey = uint8(float64(originalColor.R)*0.21 + float64(originalColor.G)*0.72 + float64(originalColor.B)*0.07)
				}
				if mode == Average {
					grey = uint8((float64(originalColor.R) + float64(originalColor.G) + float64(originalColor.B)) / 3)
				}
				if mode == Lightness {
					mmax := math.Max(float64(originalColor.R), math.Max(float64(originalColor.G), float64(originalColor.B)))
					mmin := math.Min(float64(originalColor.R), math.Min(float64(originalColor.G), float64(originalColor.B)))
					grey = uint8((mmax + mmin) / 2)
				}

				col := color.RGBA{
					grey,
					grey,
					grey,
					originalColor.A,
				}
				newImage[x][y] = col
				wg.Done()
			}(x, y)

		}
	}
	wg.Wait()
	*pixels = newImage
}
