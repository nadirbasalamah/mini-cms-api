package services

import (
	"mini-cms-api/models"
	"mini-cms-api/repositories"
)

type ContentService struct {
	Repository repositories.ContentRepository
}

func InitContentService() ContentService {
	return ContentService{
		Repository: &repositories.ContentRepositoryImpl{},
	}
}

func (cs *ContentService) GetAll() ([]models.Content, error) {
	return cs.Repository.GetAll()
}

func (cs *ContentService) GetByID(id string) (models.Content, error) {
	return cs.Repository.GetByID(id)
}

func (cs *ContentService) Create(contentReq models.ContentRequest) (models.Content, error) {
	return cs.Repository.Create(contentReq)
}

func (cs *ContentService) Update(contentReq models.ContentRequest, id string) (models.Content, error) {
	return cs.Repository.Update(contentReq, id)
}

func (cs *ContentService) Delete(id string) error {
	return cs.Repository.Delete(id)
}
