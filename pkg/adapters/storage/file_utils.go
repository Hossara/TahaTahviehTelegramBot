package storage

import (
	"fmt"
	"io"
	"path/filepath"

	"github.com/h2non/filetype"

	"taha_tahvieh_tg_bot/internal/product_storage/domain"
)

// ExtractFileName extracts the file name from a full file path.
func ExtractFileName(fullPath string) string {
	return filepath.Base(fullPath)
}

// ExtractFileFormat extracts the file format (extension)
func ExtractFileFormat(reader io.Reader) (string, error) {
	// Read the first 261 bytes (max required for detection)
	buffer := make([]byte, 261)
	n, err := reader.Read(buffer)
	if err != nil && err != io.EOF {
		return "", fmt.Errorf("error reading file: %v", err)
	}

	// Detect file type
	kind, err := filetype.Match(buffer[:n])
	if err != nil {
		return "", fmt.Errorf("unknown file type")
	}

	// Return the extension
	return kind.Extension, nil
}

// FileName generates a full file name from a name and extension.
func FileName(name, extension string) string {
	return name + "." + extension
}

func FileToFilePath(file domain.File) string {
	return file.Path + FileName(file.UUID.String(), file.Format)
}
