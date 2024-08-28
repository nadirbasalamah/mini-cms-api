package repositories_test

import (
	"errors"
	"mini-cms-api/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllContents(t *testing.T) {
	t.Run("Get All Contents | Valid", func(t *testing.T) {
		contentRepository.On("GetAll").Return([]models.Content{}, nil).Once()

		result, err := contentService.GetAll()

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("Get All Contents | Invalid", func(t *testing.T) {
		contentRepository.On("GetAll").Return([]models.Content{}, errors.New("whoops")).Once()

		result, err := contentService.GetAll()

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestGetContentByID(t *testing.T) {
	t.Run("Get Content by ID | Valid", func(t *testing.T) {
		contentRepository.On("GetByID", "1").Return(models.Content{}, nil).Once()

		result, err := contentService.GetByID("1")

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("Get Content by ID | Invalid", func(t *testing.T) {
		contentRepository.On("GetByID", "-1").Return(models.Content{}, errors.New("whoops")).Once()

		result, err := contentService.GetByID("-1")

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestCreateContent(t *testing.T) {
	t.Run("Create Content | Valid", func(t *testing.T) {
		contentRepository.On("Create", models.ContentRequest{}).Return(models.Content{}, nil).Once()

		result, err := contentService.Create(models.ContentRequest{})

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("Create Content | Invalid", func(t *testing.T) {
		contentRepository.On("Create", models.ContentRequest{}).Return(models.Content{}, errors.New("whoops")).Once()

		result, err := contentService.Create(models.ContentRequest{})

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestUpdateContent(t *testing.T) {
	t.Run("Update Content | Valid", func(t *testing.T) {
		contentRepository.On("Update", models.ContentRequest{}, "1").Return(models.Content{}, nil).Once()

		result, err := contentService.Update(models.ContentRequest{}, "1")

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("Update Content | Invalid", func(t *testing.T) {
		contentRepository.On("Update", models.ContentRequest{}, "-1").Return(models.Content{}, errors.New("whoops")).Once()

		result, err := contentService.Update(models.ContentRequest{}, "-1")

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestDeleteContent(t *testing.T) {
	t.Run("Delete Content | Valid", func(t *testing.T) {
		contentRepository.On("Delete", "1").Return(nil).Once()

		err := contentService.Delete("1")

		assert.Nil(t, err)
	})

	t.Run("Delete Content | Invalid", func(t *testing.T) {
		contentRepository.On("Delete", "-1").Return(errors.New("whoops")).Once()

		err := contentService.Delete("-1")

		assert.NotNil(t, err)
	})
}
