package categories

import (
	"mini-cms-api/businesses/categories"

	"gorm.io/gorm"
)

type categoryRepository struct {
	conn *gorm.DB
}

func NewMySQLRepository(conn *gorm.DB) categories.Repository {
	return &categoryRepository{
		conn: conn,
	}
}

func (cr *categoryRepository) GetAll() ([]categories.Domain, error) {
	var records []Category

	if err := cr.conn.Find(&records).Error; err != nil {
		return nil, err
	}

	categories := []categories.Domain{}

	for _, category := range records {
		categories = append(categories, category.ToDomain())
	}

	return categories, nil
}

func (cr *categoryRepository) GetByID(id string) (categories.Domain, error) {
	var category Category

	if err := cr.conn.First(&category, "id = ?", id).Error; err != nil {
		return categories.Domain{}, err
	}

	return category.ToDomain(), nil
}

func (cr *categoryRepository) Create(categoryReq *categories.Domain) (categories.Domain, error) {
	record := FromDomain(categoryReq)

	result := cr.conn.Create(&record)

	if err := result.Error; err != nil {
		return categories.Domain{}, err
	}

	if err := result.Last(&record).Error; err != nil {
		return categories.Domain{}, err
	}

	return record.ToDomain(), nil
}

func (cr *categoryRepository) Update(categoryReq *categories.Domain, id string) (categories.Domain, error) {
	category, err := cr.GetByID(id)

	if err != nil {
		return categories.Domain{}, err
	}

	updatedCategory := FromDomain(&category)

	updatedCategory.Name = categoryReq.Name

	if err := cr.conn.Save(&updatedCategory).Error; err != nil {
		return categories.Domain{}, err
	}

	return updatedCategory.ToDomain(), nil
}

func (cr *categoryRepository) Delete(id string) error {
	category, err := cr.GetByID(id)

	if err != nil {
		return err
	}

	deletedCategory := FromDomain(&category)

	if err := cr.conn.Delete(&deletedCategory).Error; err != nil {
		return err
	}

	return nil
}
