package main

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	_ "image/jpeg"
	"image/png"
	"log"
	"os"
	"path/filepath"
)

func MyOpen(fn string) Tensor {
	f, err := os.Open(fn)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	// to tensor
	size := img.Bounds().Size()
	var pixels Tensor
	//put pixels into two three two dimensional array
	for i := 0; i < size.X; i++ {
		var y []color.Color
		for j := 0; j < size.Y; j++ {
			y = append(y, img.At(i, j))
		}
		pixels = append(pixels, y)
	}
	return pixels
}

func MySave(fn string, pixels Tensor) {
	rect := image.Rect(0, 0, len(pixels), len(pixels[0]))
	nImg := image.NewRGBA(rect)

	for x := 0; x < len(pixels); x++ {
		for y := 0; y < len(pixels[0]); y++ {
			q := pixels[x]
			if q == nil {
				continue
			}
			p := pixels[x][y]
			if p == nil {
				continue
			}
			original, ok := color.RGBAModel.Convert(p).(color.RGBA)
			if ok {
				nImg.Set(x, y, original)
			}
		}
	}

	fg, err := os.Create(fn)
	if err != nil {
		log.Fatal(err)
	}
	defer fg.Close()
	ext := filepath.Ext(fn)
	if ext == ".jpg" {
		err = jpeg.Encode(fg, nImg, nil)
	}
	if ext == ".png" {
		err = png.Encode(fg, nImg)
	}
	if err != nil {
		fmt.Println("Encoding error", err)
	}
}
