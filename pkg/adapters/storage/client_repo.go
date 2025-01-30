package storage

import (
	"context"
	"errors"
	"fmt"
	"github.com/minio/minio-go/v7"
	"log"
	"taha_tahvieh_tg_bot/internal/product_storage/domain"
	"taha_tahvieh_tg_bot/internal/storage/port"
	minioPkg "taha_tahvieh_tg_bot/pkg/minio"
)

type storageBucket struct {
	client *minio.Client
	ctx    context.Context
}

func NewStorageRepo(config minioPkg.Config) port.ClientRepo {
	client := minioPkg.MustNewMinioClient(config)

	return &storageBucket{client: client}
}

func (s *storageBucket) UploadFile(file *domain.File) error {
	info, err := s.client.FPutObject(
		s.ctx, file.BucketName, file.UUID.String(), file.Path,
		minio.PutObjectOptions{
			ContentType: string(file.ContentType),
		},
	)

	if err == nil {
		log.Printf(
			"New file uploaded successfuly. File Name: %s\tFile Size: %d\tProduct: %d",
			FileToFileName(*file), info.Size, file.ProductID,
		)
	}

	return err
}

func (s *storageBucket) DeleteFile(bucket string, filePath string) error {
	if bucket == "" || filePath == "" {
		return errors.New("bucket and filePath cannot be empty")
	}

	return s.client.RemoveObject(s.ctx, bucket, filePath, minio.RemoveObjectOptions{})
}

func (s *storageBucket) DeleteMultipleFiles(bucket string, filePaths []string) error {
	if bucket == "" || len(filePaths) == 0 {
		return errors.New("bucket and filePaths cannot be empty")
	}

	// Create a channel to report errors
	errorCh := make(chan error)

	// Iterate through the file paths and delete each file
	for _, filePath := range filePaths {
		go func(path string) {
			err := s.client.RemoveObject(context.Background(), bucket, path, minio.RemoveObjectOptions{})
			if err != nil {
				errorCh <- fmt.Errorf("failed to delete file %s from bucket %s: %w", path, bucket, err)
				return
			}

			errorCh <- nil
		}(filePath)
	}

	// Collect errors
	var errs []error
	for range filePaths {
		if err := <-errorCh; err != nil {
			errs = append(errs, err)
		}
	}

	// Return combined errors if any
	if len(errs) > 0 {
		return fmt.Errorf("errors occurred while deleting files: %v", errs)
	}

	return nil
}

func (s *storageBucket) BucketExists(name string) (bool, error) {
	return s.client.BucketExists(s.ctx, name)
}

func (s *storageBucket) CreateBucket(name string) error {
	return s.client.MakeBucket(s.ctx, name, minio.MakeBucketOptions{})
}
