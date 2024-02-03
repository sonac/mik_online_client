package pdf

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

// List directories in a given path
func ListDirs(folderPath string) ([]string, error) {
	var folders []string // Slice to hold the paths of subdirectories

	// Walk through the folderPath
	err := filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err // Return error if any
		}
		// Check if it's a directory and not the root folder
		if info.IsDir() && path != folderPath {
			folderName := filepath.Base(path)     // Get the name of the subdirectory
			folders = append(folders, folderName) // Append the path of the subdirectory to the slice
		}
		return nil
	})

	if err != nil {
		return nil, err // Return error if encountered during walk
	}

	return folders, nil // Return the slice of subdirectories
}

func ListFiles(folderPath string) ([]string, error) {
	var folders []string // Slice to hold the paths of subdirectories

	// Walk through the folderPath
	err := filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err // Return error if any
		}
		// Check if it's a directory and not the root folder
		if !info.IsDir() && path != folderPath {
			folderName := filepath.Base(path)     // Get the name of the subdirectory
			folders = append(folders, folderName) // Append the path of the subdirectory to the slice
		}
		return nil
	})

	if err != nil {
		return nil, err // Return error if encountered during walk
	}

	return folders, nil // Return the slice of subdirectories
}

func ServePDF(w http.ResponseWriter, r *http.Request, pdfPath string) {
	// Open the file for reading
	pdfFile, err := os.Open(pdfPath)
	if err != nil {
		log.Fatal(err)
	}
	defer pdfFile.Close()

	// Get the file size
	stat, err := pdfFile.Stat()
	if err != nil {
		log.Fatal(err)
	}

	// Set the content type header
	w.Header().Set("Content-Type", "application/pdf")

	// Set the content length header
	w.Header().Set("Content-Length", strconv.FormatInt(stat.Size(), 10))

	// Send the file
	io.Copy(w, pdfFile)
}
