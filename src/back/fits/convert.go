package fits

import "fmt"

func ConvertFITSToPNG(fitsPath, pngPath string) error {
	fitsUnits, err := GetFitsUnitsFromFile(fitsPath)
	if err != nil {
		return fmt.Errorf("failed to open FITS file: %w", err)
	}

	_ = fitsUnits

	return nil
}
