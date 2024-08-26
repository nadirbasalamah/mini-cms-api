package repositories_test

import (
	"errors"
	"mini-cms-api/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {
	t.Run("GetAll | Valid", func(t *testing.T) {
		categoryRepository.On("GetAll").Return([]models.Category{}, nil).Once()

		result, err := categoryService.GetAll()

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("GetAll | Invalid", func(t *testing.T) {
		categoryRepository.On("GetAll").Return([]models.Category{}, errors.New("error")).Once()

		result, err := categoryService.GetAll()

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestGetByID(t *testing.T) {
	t.Run("GetByID | Valid", func(t *testing.T) {
		categoryRepository.On("GetByID", "1").Return(models.Category{}, nil).Once()

		result, err := categoryService.GetByID("1")

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("GetByID | Invalid", func(t *testing.T) {
		categoryRepository.On("GetByID", "0").Return(models.Category{}, errors.New("whoops")).Once()

		result, err := categoryService.GetByID("0")

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestCreate(t *testing.T) {
	t.Run("Create | Valid", func(t *testing.T) {
		categoryRepository.On("Create", models.CategoryRequest{}).Return(models.Category{}, nil).Once()

		result, err := categoryService.Create(models.CategoryRequest{})

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("Create | Invalid", func(t *testing.T) {
		categoryRepository.On("Create", models.CategoryRequest{}).Return(models.Category{}, errors.New("whoops")).Once()

		result, err := categoryService.Create(models.CategoryRequest{})

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Update | Valid", func(t *testing.T) {
		categoryRepository.On("Update", models.CategoryRequest{}, "1").Return(models.Category{}, nil).Once()

		result, err := categoryService.Update(models.CategoryRequest{}, "1")

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("Update | Invalid", func(t *testing.T) {
		categoryRepository.On("Update", models.CategoryRequest{}, "0").Return(models.Category{}, errors.New("whoops")).Once()

		result, err := categoryService.Update(models.CategoryRequest{}, "0")

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestDelete(t *testing.T) {
	t.Run("Delete | Valid", func(t *testing.T) {
		categoryRepository.On("Delete", "1").Return(nil).Once()

		err := categoryService.Delete("1")

		assert.Nil(t, err)
	})

	t.Run("Delete | Invalid", func(t *testing.T) {
		categoryRepository.On("Delete", "0").Return(errors.New("whoops")).Once()

		err := categoryService.Delete("0")

		assert.NotNil(t, err)
	})
}
