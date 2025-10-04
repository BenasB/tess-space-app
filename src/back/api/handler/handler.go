package handler

import "github.com/BenasB/tess-space-app/back/mast"

type ApiHandler struct {
	MastClient *mast.DownloadClient
}
