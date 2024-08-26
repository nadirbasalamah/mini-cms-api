package services

import (
	"mini-cms-api/models"
	"mini-cms-api/repositories"
)

type CategoryService struct {
	Repository repositories.CategoryRepository
}

func InitCategoryService() CategoryService {
	return CategoryService{
		Repository: &repositories.CategoryRepositoryImpl{},
	}
}

func (cs *CategoryService) GetAll() ([]models.Category, error) {
	return cs.Repository.GetAll()
}

func (cs *CategoryService) GetByID(id string) (models.Category, error) {
	return cs.Repository.GetByID(id)
}

func (cs *CategoryService) Create(categoryReq models.CategoryRequest) (models.Category, error) {
	return cs.Repository.Create(categoryReq)
}

func (cs *CategoryService) Update(categoryReq models.CategoryRequest, id string) (models.Category, error) {
	return cs.Repository.Update(categoryReq, id)
}

func (cs *CategoryService) Delete(id string) error {
	return cs.Repository.Delete(id)
}
