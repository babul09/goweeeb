package main

import (
	"fmt"
	"mime"
	"os"
	"path/filepath"
	"strings"
)

func readFile(path string) ([]byte, string, error) {
	const documentRoot = "public/public"

	clean := filepath.Clean(path)
	clean = strings.TrimPrefix(clean, "/")

	if strings.HasPrefix(clean, "..") {
		return nil, "", fmt.Errorf("invalid path")
	}

	fullPath := filepath.Join(documentRoot, clean)

	data, err := os.ReadFile(fullPath)
	if err != nil {
		return nil, "", err
	}

	mimeType := mime.TypeByExtension(filepath.Ext(fullPath))
	if mimeType == "" {
		mimeType = "application/octet-stream"
	}

	return data, mimeType, nil
}
