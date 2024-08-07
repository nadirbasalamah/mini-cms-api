package repositories

import "mini-cms-api/models"

type ContentRepository interface {
	GetAll() ([]models.Content, error)
	GetByID(id string) (models.Content, error)
	Create(contentReq models.ContentRequest) (models.Content, error)
	Update(contentReq models.ContentRequest, id string) (models.Content, error)
	Delete(id string) error
}

type CategoryRepository interface {
	GetAll() ([]models.Category, error)
	GetByID(id string) (models.Category, error)
	Create(categoryReq models.CategoryRequest) (models.Category, error)
	Update(categoryReq models.CategoryRequest, id string) (models.Category, error)
	Delete(id string) error
}
