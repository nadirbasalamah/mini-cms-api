package repositories

import (
	"mini-cms-api/database"
	"mini-cms-api/models"
)

type ContentRepositoryImpl struct{}

func InitContentRepository() ContentRepository {
	return &ContentRepositoryImpl{}
}

func (cr *ContentRepositoryImpl) GetAll() ([]models.Content, error) {
	var contents []models.Content

	if err := database.DB.Find(&contents).Error; err != nil {
		return []models.Content{}, err
	}

	return contents, nil
}

func (cr *ContentRepositoryImpl) GetByID(id string) (models.Content, error) {
	var content models.Content

	if err := database.DB.First(&content, "id = ?", id).Error; err != nil {
		return models.Content{}, err
	}

	return content, nil
}

func (cr *ContentRepositoryImpl) Create(contentReq models.ContentRequest) (models.Content, error) {
	var content models.Content = models.Content{
		Title:       contentReq.Title,
		Description: contentReq.Description,
	}

	result := database.DB.Create(&content)

	if err := result.Error; err != nil {
		return models.Content{}, err
	}

	if err := result.Last(&content).Error; err != nil {
		return models.Content{}, err
	}

	return content, nil
}

func (cr *ContentRepositoryImpl) Update(contentReq models.ContentRequest, id string) (models.Content, error) {
	content, err := cr.GetByID(id)

	if err != nil {
		return models.Content{}, err
	}

	content.Title = contentReq.Title
	content.Description = contentReq.Description

	if err := database.DB.Save(&content).Error; err != nil {
		return models.Content{}, err
	}

	return content, nil
}

func (cr *ContentRepositoryImpl) Delete(id string) error {
	content, err := cr.GetByID(id)

	if err != nil {
		return err
	}

	if err := database.DB.Delete(&content).Error; err != nil {
		return err
	}

	return nil
}
