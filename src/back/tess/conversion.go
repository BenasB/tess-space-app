package tess

import (
	"fmt"
	"image"

	"github.com/BenasB/tess-space-app/back/utils"
)

func ConvertFFIToPng(fitsPath, pngPath string) error {
	fitsUnits, err := utils.GetFitsUnitsFromFile(fitsPath)
	if err != nil {
		return fmt.Errorf("failed to open FITS file: %w", err)
	}

	// Calibrated image is always the second one in FFI
	img, err := utils.ConvertFitsUnitToGrayscaleImage(
		fitsUnits[1],
		2,
		99,
		&image.Rectangle{Min: image.Point{X: 44, Y: 0}, Max: image.Point{X: fitsUnits[1].Naxis[0] - 44, Y: fitsUnits[1].Naxis[1] - 30}},
	)
	if err != nil {
		return fmt.Errorf("failed to convert FITS unit to image: %w", err)
	}

	if err := utils.ExportImageToPng(img, pngPath); err != nil {
		return fmt.Errorf("failed to export image to PNG: %w", err)
	}

	return nil
}
