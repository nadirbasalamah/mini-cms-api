package repositories

import "mini-cms-api/models"

type ContentRepository interface {
	GetAll() ([]models.Content, error)
	GetByID(id string) (models.Content, error)
	Create(contentReq models.ContentRequest) (models.Content, error)
	Update(contentReq models.ContentRequest, id string) (models.Content, error)
	Delete(id string) error
}
