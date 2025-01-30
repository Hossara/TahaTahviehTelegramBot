package minio

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
)

type Config struct {
	Endpoint        string
	AccessKeyID     string
	SecretAccessKey string
	SSL             bool
}

func NewMinioClient(config Config) (*minio.Client, error) {
	return minio.New(config.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.AccessKeyID, config.SecretAccessKey, ""),
		Secure: config.SSL,
	})
}

func MustNewMinioClient(config Config) *minio.Client {
	client, err := NewMinioClient(config)

	if err != nil {
		log.Fatalf("Failed to create minio client: %v", err)
	}

	return client
}
