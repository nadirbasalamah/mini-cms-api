package services

import (
	"mini-cms-api/models"
	"mini-cms-api/repositories"
)

type ContentService struct {
	repository repositories.ContentRepository
}

func InitContentService() ContentService {
	return ContentService{
		repository: &repositories.ContentRepositoryImpl{},
	}
}

func (cs *ContentService) GetAll() ([]models.Content, error) {
	return cs.repository.GetAll()
}

func (cs *ContentService) GetByID(id string) (models.Content, error) {
	return cs.repository.GetByID(id)
}

func (cs *ContentService) Create(contentReq models.ContentRequest) (models.Content, error) {
	return cs.repository.Create(contentReq)
}

func (cs *ContentService) Update(contentReq models.ContentRequest, id string) (models.Content, error) {
	return cs.repository.Update(contentReq, id)
}

func (cs *ContentService) Delete(id string) error {
	return cs.repository.Delete(id)
}
