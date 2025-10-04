package main

import (
	"log"
	"os"
	"strconv"

	"github.com/BenasB/tess-space-app/back/tess"
	"github.com/BenasB/tess-space-app/back/utils"
)

func main() {
	if len(os.Args) < 6 {
		log.Fatalln("Usage: go run . <input1.fits> <input2.fits> <input3.fits> <input4.fits> <output.png> <downsampling factor?>")
	}

	fitPaths := os.Args[1:5]
	pngPath := os.Args[5]

	var downSamplingFactor *int
	if len(os.Args) >= 7 {
		factor, err := strconv.Atoi(os.Args[6])
		if err != nil {
			log.Fatalf("invalid downsampling factor: %v", err)
		}
		downSamplingFactor = &factor
	}

	var fitPathsArray [4]string
	for i := range 4 {
		fitPathsArray[i] = fitPaths[i]
	}

	img, err := tess.ConvertCamFFIsToImage(fitPathsArray)
	if err != nil {
		log.Fatalf("error converting cam FFIs to PNG: %v", err)
	}

	if downSamplingFactor != nil && *downSamplingFactor > 1 {
		img = utils.Downsample(img, *downSamplingFactor)
	}

	if err := utils.ExportImageToPng(img, pngPath); err != nil {
		log.Fatalf("failed to export image to PNG: %v", err)
	}

	log.Printf("Finished converting cam %s to %s\n", fitPaths, pngPath)
}
