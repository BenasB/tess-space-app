package mast

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/patrickmn/go-cache"
	"golang.org/x/sync/singleflight"
)

type DownloadClient struct {
	httpClient    *http.Client
	storage       *Storage
	downloadGroup *singleflight.Group
}

func NewDownloadClient(storage *Storage) *DownloadClient {
	return &DownloadClient{
		httpClient:    &http.Client{},
		downloadGroup: &singleflight.Group{},

		storage: storage,
	}
}

func (c *DownloadClient) DownloadFile(url string) (string, error) {
	nameParts := strings.Split(url, "/")
	fileName := nameParts[len(nameParts)-1]

	if found := c.storage.Get(fileName); found {
		return fileName, nil
	}

	_, err, _ := c.downloadGroup.Do(url, func() (any, error) {
		if err := c.download(url); err != nil {
			return nil, err
		}

		c.storage.Store(fileName, true, cache.DefaultExpiration)
		return nil, nil
	})
	if err != nil {
		return "", err
	}

	return fileName, err
}

func (c *DownloadClient) download(url string) error {
	name := strings.Split(url, "/")
	if len(name) == 0 {
		return fmt.Errorf("invalid URL: %s", url)
	}
	fileName := name[len(name)-1]
	outputPath := filepath.Join(cacheDir, fileName)

	var fileExists bool
	fileInfo, err := os.Stat(outputPath)
	if err == nil {
		fileExists = true
	} else if !os.IsNotExist(err) {
		return fmt.Errorf("could not stat output file: %w", err)
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Errorf("failed to create GET request: %v", err)
	}

	var file *os.File
	var openFlags int

	if fileExists {
		req.Header.Set("Range", fmt.Sprintf("bytes=%d-", fileInfo.Size()))
		openFlags = os.O_APPEND | os.O_WRONLY
	} else {
		req.Header.Set("Range", "bytes=0-")
		openFlags = os.O_CREATE | os.O_WRONLY
	}

	file, err = os.OpenFile(outputPath, openFlags, 0644)
	if err != nil {
		return fmt.Errorf("could not open or create file: %w", err)
	}
	defer file.Close()

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to make request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusPartialContent {
		return fmt.Errorf("failed to download file: received status code %d", resp.StatusCode)
	}

	if _, err := io.Copy(file, resp.Body); err != nil {
		return fmt.Errorf("could not write to file: %w", err)
	}

	return nil
}
