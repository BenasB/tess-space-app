package handler

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/BenasB/tess-space-app/back/tess"
	"github.com/BenasB/tess-space-app/back/utils"
)

func (h *ApiHandler) DownloadCCD(w http.ResponseWriter, r *http.Request) {
	resource := r.URL.Query().Get("resource")
	if resource == "" {
		log.Printf("Missing resource parameter")
		http.Error(w, "Missing resource parameter", http.StatusBadRequest)
		return
	} else if !strings.Contains(resource, "_ffic.fits") {
		log.Printf("Invalid resource parameter: %s", resource)
		http.Error(w, "Invalid resource parameter, must contain '_ffic.fits'", http.StatusBadRequest)
		return
	}

	localPath, err := h.MastClient.DownloadSingleFile(resource)
	if err != nil {
		log.Printf("Failed to download file: %v", err)
		http.Error(w, fmt.Sprintf("Failed to download file: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	img, err := tess.ConvertFFIToImage(localPath)
	if err != nil {
		log.Printf("Failed to convert file to image: %v", err)
		http.Error(w, fmt.Sprintf("Failed to convert file to image: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	buffer, err := utils.ExportImageToPngBuffer(img)
	if err != nil {
		log.Printf("Failed to export image to bytes buffer: %v", err)
		http.Error(w, fmt.Sprintf("Failed to export image to bytes buffer: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "image/png")

	if _, err := w.Write(buffer.Bytes()); err != nil {
		log.Printf("Failed to write image to response: %v", err)
	}

	log.Printf("Image sent successfully: %s", resource)
}
