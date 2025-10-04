package main

import (
	"log"
	"net/http"

	"github.com/BenasB/tess-space-app/back/api/router"
	"github.com/BenasB/tess-space-app/back/mast"
)

func main() {
	ms := mast.NewStorage()
	if err := ms.Start(); err != nil {
		log.Fatalf("Failed to initialize mast storage: %v", err)
	}
	mdc := mast.NewDownloadClient(ms)

	r := router.New(mdc)

	listenAddress := ":8081"
	log.Printf("Starting to listen on: %s\n", listenAddress)
	if err := http.ListenAndServe(listenAddress, r); err != nil {
		log.Fatalf("ListenAndServe failed: %v", err)
	}
}
