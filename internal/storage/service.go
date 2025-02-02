package storage

import (
	"context"
	"errors"
	"fmt"
	"io"
	"taha_tahvieh_tg_bot/internal/product_storage/domain"
	storageDomain "taha_tahvieh_tg_bot/internal/storage/domain"
	"taha_tahvieh_tg_bot/internal/storage/port"
	"taha_tahvieh_tg_bot/pkg/adapters/storage"

	productDomain "taha_tahvieh_tg_bot/internal/product/domain"
)

type service struct {
	ctx        context.Context
	repo       port.Repo
	clientRepo port.ClientRepo
}

func NewService(ctx context.Context, repo port.Repo, clientRepo port.ClientRepo) port.Service {
	return &service{
		ctx:        ctx,
		repo:       repo,
		clientRepo: clientRepo,
	}
}

func (r *service) InitBucket(name string) error {
	exists, err := r.clientRepo.BucketExists(name)

	if err != nil {
		return errors.New(fmt.Sprint("Error checking bucket existence: ", err))
	}

	if !exists {
		err := r.clientRepo.CreateBucket(name)

		if err != nil {
			return errors.New(fmt.Sprint("Error creating bucket: ", err))
		}

		return nil
	}

	return err
}

func (r *service) UploadFile(file *domain.File, url string) (storageDomain.FileID, error) {
	err := r.clientRepo.UploadFile(file, url)

	if err != nil {
		return 0, fmt.Errorf("error uploading file: %v", err)
	}

	id, err := r.repo.Insert(file)

	if err != nil {
		errDel := r.clientRepo.DeleteFile(file.BucketName, storage.FileToFilePath(*file))

		if errDel != nil {
			return 0, fmt.Errorf("error inserting & deleting file: %v", err)
		}

		return 0, fmt.Errorf("error inserting file: %v", err)
	}

	return id, nil
}

func (r *service) RemoveFile(filePath string) error {
	//TODO implement me
	panic("implement me")
}

func (r *service) RemoveAllFiles(files []domain.File) error {
	for _, file := range files {
		err := r.clientRepo.DeleteFile(file.BucketName, storage.FileToFilePath(file))

		if err != nil {
			return fmt.Errorf("error deleting file: %v", err)
		}
	}
	return nil
}

func (r *service) RemoveAllProductFiles(productID productDomain.ProductID, files []domain.File) error {
	err := r.repo.DeleteAllFilesByProductID(productID)

	if err != nil {
		return fmt.Errorf("error deleting files: %v", err)
	}

	err = r.RemoveAllFiles(files)

	if err != nil {
		return err
	}

	return nil
}

func (r *service) GetProductFile(file domain.File) (io.ReadCloser, error) {
	streamFile, err := r.clientRepo.StreamFile(file.BucketName, storage.FileToFilePath(file))

	if err != nil {
		return nil, fmt.Errorf("error streaming file: %v", err)
	}

	return streamFile, nil
}

func (r *service) RunMigrations() error {
	return r.repo.RunMigrations()
}
