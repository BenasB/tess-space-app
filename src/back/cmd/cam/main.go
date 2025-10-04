package main

import (
	"fmt"
	"log"
	"os"

	"github.com/BenasB/tess-space-app/back/tess"
)

func main() {
	if len(os.Args) < 6 {
		fmt.Println("Usage: go run . <input1.fits> <input2.fits> <input3.fits> <input4.fits> <output.png>")
		return
	}

	fitPaths := os.Args[1:5]
	pngPath := os.Args[5]

	var fitPathsArray [4]string
	copy(fitPathsArray[:], fitPaths)

	if err := tess.ConvertCamFFIsToPng(fitPathsArray, pngPath); err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Printf("Finished converting %s to %s\n", fitPaths, pngPath)
}
