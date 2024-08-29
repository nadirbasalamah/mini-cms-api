package contents_test

import (
	"errors"
	"mini-cms-api/businesses/contents"
	_contentMock "mini-cms-api/businesses/contents/mocks"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	contentRepository _contentMock.Repository
	contentUseCase    contents.UseCase
)

func TestMain(m *testing.M) {
	contentUseCase = contents.NewContentUseCase(&contentRepository)

	os.Exit(m.Run())
}

func TestGetAll(t *testing.T) {
	t.Run("GetAll | Valid", func(t *testing.T) {
		contentRepository.On("GetAll").Return([]contents.Domain{}, nil).Once()

		result, err := contentUseCase.GetAll()

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("GetAll | Invalid", func(t *testing.T) {
		contentRepository.On("GetAll").Return([]contents.Domain{}, errors.New("error")).Once()

		result, err := contentUseCase.GetAll()

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestGetByID(t *testing.T) {
	t.Run("GetByID | Valid", func(t *testing.T) {
		contentRepository.On("GetByID", "1").Return(contents.Domain{}, nil).Once()

		result, err := contentUseCase.GetByID("1")

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("GetByID | Invalid", func(t *testing.T) {
		contentRepository.On("GetByID", "0").Return(contents.Domain{}, errors.New("whoops")).Once()

		result, err := contentUseCase.GetByID("0")

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestCreate(t *testing.T) {
	t.Run("Create | Valid", func(t *testing.T) {
		contentRepository.On("Create", &contents.Domain{}).Return(contents.Domain{}, nil).Once()

		result, err := contentUseCase.Create(&contents.Domain{})

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("Create | Invalid", func(t *testing.T) {
		contentRepository.On("Create", &contents.Domain{}).Return(contents.Domain{}, errors.New("whoops")).Once()

		result, err := contentUseCase.Create(&contents.Domain{})

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Update | Valid", func(t *testing.T) {
		contentRepository.On("Update", &contents.Domain{}, "1").Return(contents.Domain{}, nil).Once()

		result, err := contentUseCase.Update(&contents.Domain{}, "1")

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("Update | Invalid", func(t *testing.T) {
		contentRepository.On("Update", &contents.Domain{}, "0").Return(contents.Domain{}, errors.New("whoops")).Once()

		result, err := contentUseCase.Update(&contents.Domain{}, "0")

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestDelete(t *testing.T) {
	t.Run("Delete | Valid", func(t *testing.T) {
		contentRepository.On("Delete", "1").Return(nil).Once()

		err := contentUseCase.Delete("1")

		assert.Nil(t, err)
	})

	t.Run("Delete | Invalid", func(t *testing.T) {
		contentRepository.On("Delete", "0").Return(errors.New("whoops")).Once()

		err := contentUseCase.Delete("0")

		assert.NotNil(t, err)
	})
}
