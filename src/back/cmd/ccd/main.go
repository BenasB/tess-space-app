package main

import (
	"log"
	"os"
	"strconv"

	"github.com/BenasB/tess-space-app/back/tess"
	"github.com/BenasB/tess-space-app/back/utils"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatalln("Usage: go run . <input.fits> <output.png> <downsampling factor?>")
	}

	fitsPath, pngPath := os.Args[1], os.Args[2]

	var downSamplingFactor *int
	if len(os.Args) >= 4 {
		factor, err := strconv.Atoi(os.Args[3])
		if err != nil {
			log.Fatalf("invalid downsampling factor: %v", err)
		}
		downSamplingFactor = &factor
	}

	img, err := tess.ConvertFFIToImage(fitsPath)
	if err != nil {
		log.Fatalf("error converting ccd FFIs to PNG: %v", err)
	}

	if downSamplingFactor != nil && *downSamplingFactor > 1 {
		img = utils.Downsample(img, *downSamplingFactor)
	}

	if err := utils.ExportImageToPng(img, pngPath); err != nil {
		log.Fatalf("failed to export image to PNG: %v", err)
	}

	log.Printf("Finished converting ccd %s to %s\n", fitsPath, pngPath)
}
