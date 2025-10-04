package fits

import (
	"fmt"
	"image"
	"image/color"
	"os"

	"github.com/BenasB/tess-space-app/back/math"
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

func ConvertFitsUnitToGrayscaleImage(imgUnit *fits.Unit, minPercentile, maxPercentile float32) (*image.Gray, error) {
	if len(imgUnit.Naxis) != 2 {
		return nil, fmt.Errorf("expected a 2D image, but got %d dimensions", len(imgUnit.Naxis))
	}

	width, height := imgUnit.Naxis[0], imgUnit.Naxis[1]
	imageData, ok := imgUnit.Data.([]float32)
	if !ok {
		return nil, fmt.Errorf("image data is not in expected format []float32")
	}

	vmin, vmax := math.GetMinMaxPercentiles(imageData, minPercentile, maxPercentile)
	img := image.NewGray(image.Rect(0, 0, width, height))
	for y := range height {
		for x := range width {
			pixelValue := imageData[y*width+x]

			if pixelValue < vmin {
				pixelValue = vmin
			} else if pixelValue > vmax {
				pixelValue = vmax
			}

			scaledValue := uint8(255 * (pixelValue - vmin) / (vmax - vmin))
			img.SetGray(x, y, color.Gray{Y: scaledValue})
		}
	}

	return img, nil
}
