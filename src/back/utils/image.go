package utils

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	_ "image/png" // Import for side-effects to register PNG format

	"github.com/anthonynsimon/bild/imgio"
	"github.com/anthonynsimon/bild/transform"
)

func ExportImageToPng(img image.Image, filePath string) error {
	if err := imgio.Save(filePath, img, imgio.PNGEncoder()); err != nil {
		return fmt.Errorf("failed to save image to %s: %w", filePath, err)
	}
	return nil
}

var zeroPoint = image.Point{0, 0}

func Tile2x2(topLeft, topRight, bottomLeft, bottomRight *image.RGBA) (*image.RGBA, error) {
	topLeftBounds := topLeft.Bounds()
	if topLeftBounds != topRight.Bounds() ||
		topLeftBounds != bottomLeft.Bounds() ||
		topLeftBounds != bottomRight.Bounds() {
		return nil, fmt.Errorf("all images must have the same dimensions")
	}

	width := topLeft.Bounds().Dx()
	height := topLeft.Bounds().Dy()

	combinedImg := image.NewRGBA(image.Rect(0, 0, width*2, height*2))

	draw.Draw(combinedImg, image.Rect(0, 0, width, height), topLeft, zeroPoint, draw.Src)
	draw.Draw(combinedImg, image.Rect(width, 0, width*2, height), topRight, zeroPoint, draw.Src)
	draw.Draw(combinedImg, image.Rect(0, height, width, height*2), bottomLeft, zeroPoint, draw.Src)
	draw.Draw(combinedImg, image.Rect(width, height, width*2, height*2), bottomRight, zeroPoint, draw.Src)

	return combinedImg, nil
}

func Stack(images ...*image.RGBA) (*image.RGBA, error) {
	if len(images) == 0 {
		return nil, fmt.Errorf("no images provided to stack")
	}

	width := images[0].Bounds().Dx()
	totalHeight := 0

	for i, img := range images {
		if img.Bounds().Dx() != width {
			return nil, fmt.Errorf("image %d has a different width: expected %d, got %d", i, width, img.Bounds().Dx())
		}
		totalHeight += img.Bounds().Dy()
	}

	stackedImg := image.NewRGBA(image.Rect(0, 0, width, totalHeight))
	currentY := 0
	for _, img := range images {
		height := img.Bounds().Dy()
		draw.Draw(stackedImg, image.Rect(0, currentY, width, currentY+height), img, zeroPoint, draw.Src)
		currentY += height
	}
	return stackedImg, nil
}

func ConvertValuesToRGBAImage(
	originalWidth, originalHeight int,
	imageData []float32,
	minValue, maxValue float32,
) (*image.RGBA, error) {
	img := image.NewRGBA(image.Rect(0, 0, originalWidth, originalHeight))
	scale := float32(0)
	if maxValue-minValue > 0 {
		scale = 255.0 / (maxValue - minValue)
	}

	for y := range originalHeight {
		for x := range originalWidth {
			sourceIdx := y*originalWidth + x
			pixelValue := imageData[sourceIdx]

			pixelValue = Clamp(pixelValue, minValue, maxValue)
			grayValue := uint8((pixelValue - minValue) * scale)
			rgbaColor := GetFalseColor(grayValue)
			destIdx := y*img.Stride + x*4

			img.Pix[destIdx+0] = rgbaColor.R
			img.Pix[destIdx+1] = rgbaColor.G
			img.Pix[destIdx+2] = rgbaColor.B
			img.Pix[destIdx+3] = rgbaColor.A
		}
	}
	return img, nil
}

func TransformRotate180(src *image.RGBA) *image.RGBA {
	return transform.Rotate(src, 180, nil)
}

func GetFalseColor(value uint8) color.RGBA {
	g := min(uint8(-0.00043*float32(value)*float32(value)+1.087*float32(value)+3.633), 255)
	b := min(uint8(-0.00080*float32(value)*float32(value)+1.119*float32(value)+17.232), 255)
	return color.RGBA{value, g, b, 255}
}
