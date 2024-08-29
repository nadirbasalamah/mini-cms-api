package categories_test

import (
	"errors"
	"mini-cms-api/businesses/categories"
	_categoryMock "mini-cms-api/businesses/categories/mocks"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	categoryRepository _categoryMock.Repository
	categoryUseCase    categories.UseCase
)

func TestMain(m *testing.M) {
	categoryUseCase = categories.NewCategoryUseCase(&categoryRepository)

	os.Exit(m.Run())
}

func TestGetAll(t *testing.T) {
	t.Run("GetAll | Valid", func(t *testing.T) {
		categoryRepository.On("GetAll").Return([]categories.Domain{}, nil).Once()

		result, err := categoryUseCase.GetAll()

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("GetAll | Invalid", func(t *testing.T) {
		categoryRepository.On("GetAll").Return([]categories.Domain{}, errors.New("error")).Once()

		result, err := categoryUseCase.GetAll()

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestGetByID(t *testing.T) {
	t.Run("GetByID | Valid", func(t *testing.T) {
		categoryRepository.On("GetByID", "1").Return(categories.Domain{}, nil).Once()

		result, err := categoryUseCase.GetByID("1")

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("GetByID | Invalid", func(t *testing.T) {
		categoryRepository.On("GetByID", "0").Return(categories.Domain{}, errors.New("whoops")).Once()

		result, err := categoryUseCase.GetByID("0")

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestCreate(t *testing.T) {
	t.Run("Create | Valid", func(t *testing.T) {
		categoryRepository.On("Create", &categories.Domain{}).Return(categories.Domain{}, nil).Once()

		result, err := categoryUseCase.Create(&categories.Domain{})

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("Create | Invalid", func(t *testing.T) {
		categoryRepository.On("Create", &categories.Domain{}).Return(categories.Domain{}, errors.New("whoops")).Once()

		result, err := categoryUseCase.Create(&categories.Domain{})

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Update | Valid", func(t *testing.T) {
		categoryRepository.On("Update", &categories.Domain{}, "1").Return(categories.Domain{}, nil).Once()

		result, err := categoryUseCase.Update(&categories.Domain{}, "1")

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("Update | Invalid", func(t *testing.T) {
		categoryRepository.On("Update", &categories.Domain{}, "0").Return(categories.Domain{}, errors.New("whoops")).Once()

		result, err := categoryUseCase.Update(&categories.Domain{}, "0")

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestDelete(t *testing.T) {
	t.Run("Delete | Valid", func(t *testing.T) {
		categoryRepository.On("Delete", "1").Return(nil).Once()

		err := categoryUseCase.Delete("1")

		assert.Nil(t, err)
	})

	t.Run("Delete | Invalid", func(t *testing.T) {
		categoryRepository.On("Delete", "0").Return(errors.New("whoops")).Once()

		err := categoryUseCase.Delete("0")

		assert.NotNil(t, err)
	})
}
