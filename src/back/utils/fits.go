package utils

import (
	"fmt"
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

func MapFitsImageUnitToImageValues(fitsUnits *fits.Unit) (imageWidth, imageHeight int, imageData []float32, err error) {
	if fitsUnits == nil {
		return 0, 0, nil, fmt.Errorf("fits unit is nil")
	}

	if len(fitsUnits.Naxis) < 2 {
		return 0, 0, nil, fmt.Errorf("fits image unit is expected to have at least 2 dimensions")
	}
	imageWidth, imageHeight = fitsUnits.Naxis[0], fitsUnits.Naxis[1]
	imageData, ok := fitsUnits.Data.([]float32)
	if !ok {
		return 0, 0, nil, fmt.Errorf("image data is not in expected format []float32")
	}
	return imageWidth, imageHeight, imageData, nil
}
