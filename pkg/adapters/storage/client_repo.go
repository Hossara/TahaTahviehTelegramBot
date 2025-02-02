package storage

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/minio/minio-go/v7"
	"golang.org/x/net/proxy"
	"io"
	"log"
	"net"
	"net/http"
	"taha_tahvieh_tg_bot/config"
	"taha_tahvieh_tg_bot/internal/product_storage/domain"
	storageDomain "taha_tahvieh_tg_bot/internal/storage/domain"
	"taha_tahvieh_tg_bot/internal/storage/port"
	minioPkg "taha_tahvieh_tg_bot/pkg/minio"
)

type storageBucket struct {
	client *minio.Client
	ctx    context.Context
	svcCfg config.ServerConfig
}

func NewStorageRepo(svcCfg config.ServerConfig, config minioPkg.Config, ctx context.Context) port.ClientRepo {
	client := minioPkg.MustNewMinioClient(config)

	return &storageBucket{client: client, ctx: ctx, svcCfg: svcCfg}
}

func (s *storageBucket) StreamFile(bucket, name string) (io.ReadCloser, error) {
	object, err := s.client.GetObject(s.ctx, bucket, name, minio.GetObjectOptions{})

	if err != nil {
		return nil, err
	}

	return object, nil
}

func (s *storageBucket) getSocks() (*http.Client, error) {
	dialer, err := proxy.SOCKS5("tcp", s.svcCfg.Proxy, nil, proxy.Direct)

	if err != nil {
		return nil, fmt.Errorf("failed to create SOCKS5 dialer: %v", err)
	}

	transport := &http.Transport{
		DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
			return dialer.Dial(network, addr)
		},
	}

	client := &http.Client{Transport: transport}

	return client, err
}

func (s *storageBucket) UploadFile(file *domain.File, url string) error {
	var resp *http.Response
	var err error

	if s.svcCfg.Proxy != "" {
		client, err := s.getSocks()
		if err != nil {
			return fmt.Errorf("failed to create SOCKS5 proxy client: %v", err)
		}

		resp, err = client.Get(url)
	} else {
		resp, err = http.Get(url)
	}

	if err != nil {
		return fmt.Errorf("failed to download file: %v", err)
	}

	if resp == nil {
		return errors.New("failed to download file")
	}

	defer resp.Body.Close()

	// Read the file into memory
	fileData, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read file: %v", err)
	}

	file.ContentType = storageDomain.ContentType(http.DetectContentType(fileData))
	format, err := ExtractFileFormat(bytes.NewReader(fileData))

	if err != nil {
		return fmt.Errorf("failed to extract file format: %v", err)
	}

	file.Format = format

	info, err := s.client.PutObject(
		s.ctx, file.BucketName, FileToFilePath(*file), bytes.NewReader(fileData),
		int64(len(fileData)),
		minio.PutObjectOptions{
			ContentType: string(file.ContentType),
		},
	)

	if err == nil {
		log.Printf(
			"New file uploaded successfuly. File Name: %s\tFile Size: %d\tProduct: %d",
			FileToFilePath(*file), info.Size, file.ProductID,
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
