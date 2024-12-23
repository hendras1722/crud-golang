package helpers

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

// SaveFile uploads a file and returns the saved file path or an error
func SaveFile(file multipart.File, fileHeader *multipart.FileHeader, uploadDir string) (string, error) {
	// Ensure the upload directory exists
	err := os.MkdirAll(uploadDir, os.ModePerm)
	if err != nil {
		return "", fmt.Errorf("failed to create upload directory: %v", err)
	}

	// Build file path
	filePath := filepath.Join(uploadDir, fileHeader.Filename)

	// Create a new file in the upload directory
	outFile, err := os.Create(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to create file: %v", err)
	}
	defer outFile.Close()

	// Copy the file content
	_, err = io.Copy(outFile, file)
	if err != nil {
		return "", fmt.Errorf("failed to write file to disk: %v", err)
	}

	return filePath, nil
}
