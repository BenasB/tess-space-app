package main

import (
	"log"
	"os"
	"strconv"

	"github.com/BenasB/tess-space-app/back/tess"
	"github.com/BenasB/tess-space-app/back/utils"
)

func main() {
	if len(os.Args) < 18 {
		log.Fatalln("Usage: go run . <input11.fits> <input12.fits> <input13.fits> <input14.fits> <input21.fits> <input22.fits> <input23.fits> <input24.fits> <input31.fits> <input32.fits> <input33.fits> <input34.fits> <input41.fits> <input42.fits> <input43.fits> <input44.fits> <output.png> <downsampling factor?>")
	}

	fitPaths := os.Args[1:17]
	pngPath := os.Args[17]
	var downSamplingFactor *int
	if len(os.Args) >= 19 {
		factor, err := strconv.Atoi(os.Args[18])
		if err != nil {
			log.Fatalf("invalid downsampling factor: %v", err)
		}
		downSamplingFactor = &factor
	}

	var fitPaths2DArray [4][4]string
	for i := range 4 {
		for j := range 4 {
			fitPaths2DArray[i][j] = fitPaths[i*4+j]
		}
	}

	img, err := tess.ConvertSectorFFIsToImage(fitPaths2DArray)
	if err != nil {
		log.Fatalf("error converting sector FFIs to PNG: %v", err)
	}

	if downSamplingFactor != nil && *downSamplingFactor > 1 {
		img = utils.Downsample(img, *downSamplingFactor)
	}

	if err := utils.ExportImageToPng(img, pngPath); err != nil {
		log.Fatalf("failed to export image to PNG: %v", err)
	}

	log.Printf("Finished converting sector %s to %s\n", fitPaths, pngPath)
}
