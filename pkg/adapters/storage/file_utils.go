package storage

import (
	"path/filepath"
	"strings"
	"taha_tahvieh_tg_bot/internal/product_storage/domain"
)

// ExtractFileName extracts the file name from a full file path.
func ExtractFileName(fullPath string) string {
	return filepath.Base(fullPath)
}

// ExtractFileFormat extracts the file format (extension) from a full file path.
func ExtractFileFormat(fullPath string) string {
	return strings.TrimPrefix(filepath.Ext(fullPath), ".")
}

// FileName generates a full file name from a name and extension.
func FileName(name, extension string) string {
	return name + "." + extension
}

func FileToFileName(file domain.File) string {
	return FileName(file.UUID.String(), file.Format)
}
