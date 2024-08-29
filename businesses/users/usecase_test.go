package users_test

import (
	"errors"
	"mini-cms-api/app/middlewares"
	"mini-cms-api/businesses/users"
	_userMock "mini-cms-api/businesses/users/mocks"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	userRepository _userMock.Repository
	userUseCase    users.UseCase
)

func TestMain(m *testing.M) {
	userUseCase = users.NewUserUseCase(&userRepository, &middlewares.JWTConfig{})

	os.Exit(m.Run())
}

func TestRegister(t *testing.T) {
	t.Run("Register | Valid", func(t *testing.T) {
		userRepository.On("Register", &users.Domain{}).Return(users.Domain{}, nil).Once()

		result, err := userUseCase.Register(&users.Domain{})

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("Register | Invalid", func(t *testing.T) {
		userRepository.On("Register", &users.Domain{}).Return(users.Domain{}, errors.New("something went wrong")).Once()

		result, err := userUseCase.Register(&users.Domain{})

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestGetUserByEmail(t *testing.T) {
	t.Run("GetUserByEmail | Valid", func(t *testing.T) {
		userRepository.On("GetUserByEmail", &users.Domain{}).Return(users.Domain{}, nil).Once()

		result, err := userUseCase.Login(&users.Domain{})

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("GetUserByEmail | Invalid", func(t *testing.T) {
		userRepository.On("GetUserByEmail", &users.Domain{}).Return(users.Domain{}, errors.New("something went wrong")).Once()

		result, err := userUseCase.Login(&users.Domain{})

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestGetUserInfo(t *testing.T) {
	t.Run("GetUserInfo | Valid", func(t *testing.T) {
		userRepository.On("GetUserInfo", "1").Return(users.Domain{}, nil).Once()

		result, err := userUseCase.GetUserInfo("1")

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("GetUserInfo | Invalid", func(t *testing.T) {
		userRepository.On("GetUserInfo", "0").Return(users.Domain{}, errors.New("something went wrong")).Once()

		result, err := userUseCase.GetUserInfo("0")

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}
