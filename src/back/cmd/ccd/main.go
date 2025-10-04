package main

import (
	"fmt"
	"log"
	"os"

	"github.com/BenasB/tess-space-app/back/tess"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run . <input.fits> <output.png>")
		return
	}

	fitsPath, pngPath := os.Args[1], os.Args[2]
	if err := tess.ConvertFFIToPng(fitsPath, pngPath); err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Printf("Finished converting %s to %s\n", fitsPath, pngPath)
}
