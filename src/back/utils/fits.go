package utils

import (
	"fmt"
	"image"
	"os"

	"github.com/siravan/fits"
)

func GetFitsUnitsFromFile(filePath string) ([]*fits.Unit, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	f, err := fits.Open(file)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func ConvertFitsUnitToGrayscaleImage(
	imgUnit *fits.Unit,
	minPercentile, maxPercentile float32,
	cropRect *image.Rectangle, // Can be nil
) (*image.Gray, error) {
	if len(imgUnit.Naxis) != 2 {
		return nil, fmt.Errorf("expected a 2D image, but got %d dimensions", len(imgUnit.Naxis))
	}

	originalWidth, originalHeight := imgUnit.Naxis[0], imgUnit.Naxis[1]
	imageData, ok := imgUnit.Data.([]float32)
	if !ok {
		return nil, fmt.Errorf("image data is not in expected format []float32")
	}

	var finalRect image.Rectangle
	if cropRect != nil {
		finalRect = *cropRect
	} else {
		finalRect = image.Rect(0, 0, originalWidth, originalHeight)
	}

	finalWidth, finalHeight := finalRect.Dx(), finalRect.Dy()
	if finalWidth <= 0 || finalHeight <= 0 {
		return nil, fmt.Errorf("cropping results in a non-positive image dimension: new size %dx%d", finalWidth, finalHeight)
	}

	vmin, vmax := GetMinMaxPercentiles(imageData, minPercentile, maxPercentile)
	img := image.NewGray(image.Rect(0, 0, finalWidth, finalHeight))

	scale := float32(0)
	if vmax-vmin > 0 {
		scale = 255.0 / (vmax - vmin)
	}

	for y := range finalHeight {
		for x := range finalWidth {
			originalX := x + finalRect.Min.X
			originalY := y + finalRect.Min.Y
			pixelValue := imageData[originalY*originalWidth+originalX]
			pixelValue = Clamp(pixelValue, vmin, vmax)
			scaledValue := uint8((pixelValue - vmin) * scale)
			img.Pix[y*img.Stride+x] = scaledValue
		}
	}

	return img, nil
}
