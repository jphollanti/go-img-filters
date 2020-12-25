package main

import (
	"image/color"
	"os"
	"path/filepath"

	"gonum.org/v1/gonum/mat"
)

type Tensor = [][]color.Color

func upsideDown(pixels Tensor) {
	for i := 0; i < len(pixels); i++ {
		tr := pixels[i]
		for j := 0; j < len(tr)/2; j++ {
			k := len(tr) - j - 1
			tr[j], tr[k] = tr[k], tr[j]
		}
	}
}

func main() {
	newpath := filepath.Join(".", "target")
	os.MkdirAll(newpath, os.ModePerm)

	//tensor := MyOpen("nc2.jpg")
	//upsideDown(tensor)
	//grayScale(&tensor, Lightness)
	//MySave("target/nc2_gray_Lightness.jpg", tensor)

	//xLen := len(tensor)
	//yLen := len(tensor[0])
	//diagonal := int(math.Sqrt(float64(xLen*xLen) + float64(yLen*yLen)))
	//simpleRotationByAngle(math.Pi/10, &tensor, diagonal)
	//MySave("target/nc2_rotate_10.jpg", tensor)

	tensor := MyOpen("nc2_gray_luminosity.jpg")
	//blur := 1.0 / 9
	//boxKernel := mat.NewDense(3, 3, []float64{
	//	blur, blur, blur,
	//	blur, blur, blur,
	//	blur, blur, blur,
	//})
	//spatialFilter(&tensor, boxKernel)
	gaussianKernel := mat.NewDense(5, 5, []float64{
		1.0 / 256, 4.0 / 256, 6.0 / 256, 4.0 / 256, 1.0 / 256,
		4.0 / 256, 16.0 / 256, 24.0 / 256, 16.0 / 256, 4.0 / 256,
		6.0 / 256, 24.0 / 256, 36.0 / 256, 24.0 / 256, 6.0 / 256,
		4.0 / 256, 16.0 / 256, 24.0 / 256, 16.0 / 256, 4.0 / 256,
		1.0 / 256, 4.0 / 256, 6.0 / 256, 4.0 / 256, 1.0 / 256,
	})
	spatialFilter(&tensor, gaussianKernel)

	MySave("target/nc2_gaussianblur.jpg", tensor)
}
