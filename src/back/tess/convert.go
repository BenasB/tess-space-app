package tess

import (
	"fmt"
	"image"

	"github.com/BenasB/tess-space-app/back/fits"
	"github.com/BenasB/tess-space-app/back/imaging"
)

func ConvertFFIToPng(fitsPath, pngPath string) error {
	fitsUnits, err := fits.GetFitsUnitsFromFile(fitsPath)
	if err != nil {
		return fmt.Errorf("failed to open FITS file: %w", err)
	}

	// Calibrated image is always the second one in FFI
	img, err := fits.ConvertFitsUnitToGrayscaleImage(fitsUnits[1], 2, 99)
	if err != nil {
		return fmt.Errorf("failed to convert FITS unit to image: %w", err)
	}

	bounds := img.Bounds()
	// We need to trim top 30 rows and 44 columns from left and right
	cropRect := image.Rectangle{Min: bounds.Min.Add(image.Point{X: 44, Y: 0}), Max: bounds.Max.Sub(image.Point{X: 44, Y: 30})}
	croppedImage := img.SubImage(cropRect)

	if err := imaging.ExportImageToPng(croppedImage, pngPath); err != nil {
		return fmt.Errorf("failed to export image to PNG: %w", err)
	}

	return nil
}
