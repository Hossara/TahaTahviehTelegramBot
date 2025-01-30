package domain

type ContentType string
type FileID int64

const (
	ContentTypeOctetStream ContentType = "application/octet-stream"
	ContentTypePdf         ContentType = "application/pdf"
	ContentTypePng         ContentType = "image/png"
	ContentTypeJpeg        ContentType = "image/jpeg"
	ContentTypeCsv         ContentType = "text/csv"
)
