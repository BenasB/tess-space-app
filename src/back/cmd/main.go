package main

import (
	"fmt"
	"log"
	"os"

	"github.com/BenasB/tess-space-app/back/fits"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run . <input.fits> <output.png>")
		return
	}

	fitsPath := os.Args[1]
	pngPath := os.Args[2]

	if err := fits.ConvertFITSToPNG(fitsPath, pngPath); err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Printf("Finished converting %s to %s\n", fitsPath, pngPath)
}
