package utils

import (
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"os"
)

func ExportImageToPng(image image.Image, filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create new file %s: %w", filePath, err)
	}

	defer file.Close()
	err = png.Encode(file, image)
	if err != nil {
		return fmt.Errorf("failed to encode image to PNG: %w", err)
	}

	return nil

}

func Tile2x2(topLeft, topRight, bottomLeft, bottomRight image.Image) (image.Image, error) {
	topLeftBounds := topLeft.Bounds()
	if topLeftBounds != topRight.Bounds() ||
		topLeftBounds != bottomLeft.Bounds() ||
		topLeftBounds != bottomLeft.Bounds() {
		return nil, fmt.Errorf("all images must have the same dimensions")
	}

	width := topLeft.Bounds().Dx()
	height := topLeft.Bounds().Dy()

	combinedImg := image.NewGray(image.Rect(0, 0, width*2, height*2))

	topLeftRect := image.Rect(0, 0, width, height)
	topRightRect := image.Rect(width, 0, width*2, height)
	bottomLeftRect := image.Rect(0, height, width, height*2)
	bottomRightRect := image.Rect(width, height, width*2, height*2)

	draw.Draw(combinedImg, topLeftRect, topLeft, image.Point{0, 0}, draw.Src)
	draw.Draw(combinedImg, topRightRect, topRight, image.Point{0, 0}, draw.Src)
	draw.Draw(combinedImg, bottomLeftRect, bottomLeft, image.Point{0, 0}, draw.Src)
	draw.Draw(combinedImg, bottomRightRect, bottomRight, image.Point{0, 0}, draw.Src)

	return combinedImg, nil
}

func ConvertValuesToGrayscaleImage(
	originalWidth, originalHeight int,
	imageData []float32,
	minValue, maxValue float32,
) (*image.Gray, error) {
	img := image.NewGray(image.Rect(0, 0, originalWidth, originalHeight))
	scale := float32(0)
	if maxValue-minValue > 0 {
		scale = 255.0 / (maxValue - minValue)
	}

	for y := range originalHeight {
		for x := range originalWidth {
			pixelValue := imageData[y*originalWidth+x]
			pixelValue = Clamp(pixelValue, minValue, maxValue)
			scaledValue := uint8((pixelValue - minValue) * scale)
			img.Pix[y*img.Stride+x] = scaledValue
		}
	}

	return img, nil
}

func TransformRotate180(src *image.Gray) *image.Gray {
	bounds := src.Bounds()
	dst := image.NewGray(bounds)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			dst.Set(bounds.Max.X-x-1, bounds.Max.Y-y-1, src.At(x, y))
		}
	}
	return dst
}
