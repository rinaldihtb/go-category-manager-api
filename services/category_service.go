package services

import (
	"category-manager-api/models"
	"category-manager-api/repositories"
)

type CategoryService struct {
	repo *repositories.CategoryRepository
}

func NewCategoryService(repo *repositories.CategoryRepository) *CategoryService {
	return &CategoryService{repo: repo}
}

func (s *CategoryService) GetAll() ([]models.Category, error) {
	return s.repo.GetAll()
}

func (s *CategoryService) Create(data *models.Category) error {
	return s.repo.Create(data)
}
func (s *CategoryService) GetById(id int) (*models.Category, error) {
	return s.repo.GetById(id)
}
func (s *CategoryService) Update(category *models.Category) error {
	return s.repo.Update(category)
}
func (s *CategoryService) DeleteById(id int) error {
	return s.repo.DeleteById(id)
}
