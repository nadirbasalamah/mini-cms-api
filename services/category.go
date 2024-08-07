package services

import (
	"mini-cms-api/models"
	"mini-cms-api/repositories"
)

type CategoryService struct {
	repository repositories.CategoryRepository
}

func InitCategoryService() CategoryService {
	return CategoryService{
		repository: &repositories.CategoryRepositoryImpl{},
	}
}

func (cs *CategoryService) GetAll() ([]models.Category, error) {
	return cs.repository.GetAll()
}

func (cs *CategoryService) GetByID(id string) (models.Category, error) {
	return cs.repository.GetByID(id)
}

func (cs *CategoryService) Create(categoryReq models.CategoryRequest) (models.Category, error) {
	return cs.repository.Create(categoryReq)
}

func (cs *CategoryService) Update(categoryReq models.CategoryRequest, id string) (models.Category, error) {
	return cs.repository.Update(categoryReq, id)
}

func (cs *CategoryService) Delete(id string) error {
	return cs.repository.Delete(id)
}
