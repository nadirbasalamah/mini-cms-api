package repositories

import (
	"mini-cms-api/database"
	"mini-cms-api/models"
)

type CategoryRepositoryImpl struct{}

func InitCategoryRepository() CategoryRepository {
	return &CategoryRepositoryImpl{}
}

func (cr *CategoryRepositoryImpl) GetAll() ([]models.Category, error) {
	var categories []models.Category

	if err := database.DB.Find(&categories).Error; err != nil {
		return []models.Category{}, err
	}

	return categories, nil
}

func (cr *CategoryRepositoryImpl) GetByID(id string) (models.Category, error) {
	var category models.Category

	if err := database.DB.First(&category, "id = ?", id).Error; err != nil {
		return models.Category{}, err
	}

	return category, nil
}

func (cr *CategoryRepositoryImpl) Create(categoryReq models.CategoryRequest) (models.Category, error) {
	var category models.Category = models.Category{
		Name: categoryReq.Name,
	}

	result := database.DB.Create(&category)

	if err := result.Error; err != nil {
		return models.Category{}, err
	}

	if err := result.Last(&category).Error; err != nil {
		return models.Category{}, err
	}

	return category, nil
}

func (cr *CategoryRepositoryImpl) Update(categoryReq models.CategoryRequest, id string) (models.Category, error) {
	category, err := cr.GetByID(id)

	if err != nil {
		return models.Category{}, err
	}

	category.Name = categoryReq.Name

	if err := database.DB.Save(&category).Error; err != nil {
		return models.Category{}, err
	}

	return category, nil
}

func (cr *CategoryRepositoryImpl) Delete(id string) error {
	category, err := cr.GetByID(id)

	if err != nil {
		return err
	}

	if err := database.DB.Delete(&category).Error; err != nil {
		return err
	}

	return nil
}
