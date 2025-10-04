package tess

import (
	"fmt"
	"image"

	"github.com/BenasB/tess-space-app/back/utils"
)

const minValue = 0   // about min 2 percentile in cam4 ccd1 t0
const maxValue = 800 // about max 99 percentile in cam4 ccd1 t0

func ConvertFFIToImage(fitsPath string) (*image.Gray, error) {
	fitsUnits, err := utils.GetFitsUnitsFromFile(fitsPath)
	if err != nil || len(fitsUnits) < 2 {
		return nil, fmt.Errorf("failed to open FITS file: %w", err)
	}

	// Calibrated image is always the second one in FFI
	width, height, imageData, err := utils.MapFitsImageUnitToImageValues(fitsUnits[1])
	if err != nil {
		return nil, fmt.Errorf("failed to map FITS unit to image values: %w", err)
	}

	return utils.ConvertValuesToGrayscaleImage(
		width,
		height,
		imageData,
		minValue,
		maxValue,
	)
}

func ConvertCamFFIsToImage(fitsPaths [4]string) (*image.Gray, error) {
	ccds := utils.MapFiltered(fitsPaths[:], func(fitsPath string) (*image.Gray, bool) {
		ccd, err := ConvertFFIToImage(fitsPath)
		if err != nil {
			return nil, false
		}
		return ccd, true
	})

	if len(ccds) != 4 {
		return nil, fmt.Errorf("expected 4 valid converted ccds, but got %d", len(ccds))
	}

	return utils.Tile2x2(
		ccds[2],
		ccds[3],
		utils.TransformRotate180(ccds[1]),
		utils.TransformRotate180(ccds[0]),
	)
}

func ConvertSectorFFIsToImage(fitsPaths [4][4]string) (*image.Gray, error) {
	cams := utils.MapFiltered(fitsPaths[:], func(fitsPaths [4]string) (*image.Gray, bool) {
		cam, err := ConvertCamFFIsToImage(fitsPaths)
		if err != nil {
			return nil, false
		}
		return cam, true
	})
	if len(cams) != 4 {
		return nil, fmt.Errorf("expected 4 valid converted cams, but got %d", len(cams))
	}

	return utils.Stack(
		cams[0],
		cams[1],
		utils.TransformRotate180(cams[2]),
		utils.TransformRotate180(cams[3]),
	)
}
