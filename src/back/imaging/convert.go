package imaging

import (
	"fmt"
	"image"
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
