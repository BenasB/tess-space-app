package tess

import (
	"fmt"
	"image"

	"github.com/BenasB/tess-space-app/back/utils"
	"github.com/siravan/fits"
)

const minValue = 0   // about min 2 percentile in cam4 ccd1 t0
const maxValue = 800 // about max 99 percentile in cam4 ccd1 t0

func ConvertFFIToPng(fitsPath, pngPath string) error {
	fitsUnits, err := utils.GetFitsUnitsFromFile(fitsPath)
	if err != nil {
		return fmt.Errorf("failed to open FITS file: %w", err)
	}

	// Calibrated image is always the second one in FFI
	imageWidth, imageHeight, imageData, err := utils.MapFitsImageUnitToImageValues(fitsUnits[1])
	if err != nil {
		return fmt.Errorf("failed to map FITS unit to image values: %w", err)
	}

	img, err := utils.ConvertValuesToGrayscaleImage(
		imageWidth,
		imageHeight,
		imageData,
		minValue,
		maxValue,
	)
	if err != nil {
		return fmt.Errorf("failed to convert image values to grayscale image: %w", err)
	}

	if err := utils.ExportImageToPng(img, pngPath); err != nil {
		return fmt.Errorf("failed to export image to PNG: %w", err)
	}

	return nil
}

func ConvertCamFFIsToPng(fitsPaths [4]string, pngPath string) error {
	fitsImages := utils.MapFiltered(fitsPaths[:], func(_ int, fitsPath string) (*fits.Unit, bool) {
		units, err := utils.GetFitsUnitsFromFile(fitsPath)
		if err != nil || len(units) < 2 {
			return nil, false
		}
		// Calibrated image is always the second one in FFI
		return units[1], true
	})
	if len(fitsImages) != 4 {
		return fmt.Errorf("expected 4 valid FITS images, but got %d", len(fitsImages))
	}

	imgs := utils.MapFiltered(fitsImages, func(index int, unit *fits.Unit) (*image.Gray, bool) {
		width, height, imageData, err := utils.MapFitsImageUnitToImageValues(unit)
		if err != nil {
			return nil, false
		}

		img, err := utils.ConvertValuesToGrayscaleImage(
			width,
			height,
			imageData,
			minValue,
			maxValue,
		)
		if err != nil {
			return nil, false
		}
		return img, true
	})
	if len(imgs) != 4 {
		return fmt.Errorf("expected 4 valid converted images, but got %d", len(imgs))
	}

	img, err := utils.Tile2x2(
		imgs[2],
		imgs[3],
		utils.TransformRotate180(imgs[1]),
		utils.TransformRotate180(imgs[0]),
	)
	if err != nil {
		return fmt.Errorf("failed to tile images: %w", err)
	}

	if err := utils.ExportImageToPng(img, pngPath); err != nil {
		return fmt.Errorf("failed to export image to PNG: %w", err)
	}

	return nil
}
