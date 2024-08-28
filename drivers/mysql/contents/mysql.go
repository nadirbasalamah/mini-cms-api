package contents

import (
	"mini-cms-api/businesses/contents"

	"gorm.io/gorm"
)

type contentRepository struct {
	conn *gorm.DB
}

func NewMySQLRepository(conn *gorm.DB) contents.Repository {
	return &contentRepository{}
}

func (cr *contentRepository) GetAll() ([]contents.Domain, error) {
	var records []Content

	if err := cr.conn.Preload("Category").Find(&records).Error; err != nil {
		return nil, err
	}

	contents := []contents.Domain{}

	for _, content := range records {
		contents = append(contents, content.ToDomain())
	}

	return contents, nil
}

func (cr *contentRepository) GetByID(id string) (contents.Domain, error) {
	var content Content

	if err := cr.conn.Preload("Category").First(&content, "id = ?", id).Error; err != nil {
		return contents.Domain{}, err
	}

	return content.ToDomain(), nil
}

func (cr *contentRepository) Create(contentReq *contents.Domain) (contents.Domain, error) {
	record := FromDomain(contentReq)

	result := cr.conn.Create(&record)

	if err := result.Error; err != nil {
		return contents.Domain{}, err
	}

	if err := result.Last(&record).Error; err != nil {
		return contents.Domain{}, err
	}

	return record.ToDomain(), nil
}

func (cr *contentRepository) Update(contentReq *contents.Domain, id string) (contents.Domain, error) {
	content, err := cr.GetByID(id)

	if err != nil {
		return contents.Domain{}, err
	}

	updatedContent := FromDomain(&content)

	updatedContent.Title = contentReq.Title
	updatedContent.Description = contentReq.Description
	updatedContent.CategoryID = contentReq.CategoryID

	if err := cr.conn.Save(&updatedContent).Error; err != nil {
		return contents.Domain{}, err
	}

	return updatedContent.ToDomain(), nil
}

func (cr *contentRepository) Delete(id string) error {
	content, err := cr.GetByID(id)

	if err != nil {
		return err
	}

	deletedContent := FromDomain(&content)

	if err := cr.conn.Delete(&deletedContent).Error; err != nil {
		return err
	}

	return nil
}
