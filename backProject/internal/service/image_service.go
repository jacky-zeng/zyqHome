package service

import (
	"errors"
	"os"
	"path/filepath"

	"zyqHome/backProject/internal/model"
	"zyqHome/backProject/internal/repository"
)

type ImageService struct {
	repo      *repository.ImageRepo
	uploadDir string
}

func NewImageService(repo *repository.ImageRepo, uploadDir string) *ImageService {
	return &ImageService{repo: repo, uploadDir: uploadDir}
}

func (s *ImageService) GetList(category string, page, pageSize int) ([]model.Image, int64, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}
	return s.repo.FindAll(category, page, pageSize)
}

func (s *ImageService) GetUserList(userID uint, category string, page, pageSize int) ([]model.Image, int64, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}
	return s.repo.FindAllByUser(userID, category, page, pageSize)
}

func (s *ImageService) GetCategories() ([]string, error) {
	return s.repo.FindCategories()
}

func (s *ImageService) GetUserCategories(userID uint) ([]string, error) {
	return s.repo.FindCategoriesByUser(userID)
}

func (s *ImageService) GetByID(id uint) (*model.Image, error) {
	return s.repo.FindByID(id)
}

func (s *ImageService) Create(image *model.Image) error {
	return s.repo.Create(image)
}

func (s *ImageService) Update(image *model.Image) error {
	return s.repo.Update(image)
}

func (s *ImageService) Delete(id uint) error {
	img, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	// Delete the physical file
	filename := filepath.Base(img.URL)
	filePath := filepath.Join(s.uploadDir, filename)
	os.Remove(filePath)

	return s.repo.Delete(id)
}

func (s *ImageService) DeleteByUser(userID, imageID uint) error {
	img, err := s.repo.FindByID(imageID)
	if err != nil {
		return err
	}
	if img.UserID != userID {
		return errors.New("无权删除该图片")
	}

	filename := filepath.Base(img.URL)
	filePath := filepath.Join(s.uploadDir, filename)
	os.Remove(filePath)

	return s.repo.Delete(imageID)
}
