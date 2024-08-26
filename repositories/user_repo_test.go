package repositories_test

import (
	"errors"
	"mini-cms-api/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	t.Run("Register | Valid", func(t *testing.T) {
		userRepository.On("Register", models.RegisterRequest{}).Return(models.User{}, nil).Once()

		result, err := userService.Register(models.RegisterRequest{})

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("Register | Invalid", func(t *testing.T) {
		userRepository.On("Register", models.RegisterRequest{}).Return(models.User{}, errors.New("something went wrong")).Once()

		result, err := userService.Register(models.RegisterRequest{})

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestGetByEmail(t *testing.T) {
	t.Run("GetByEmail | Valid", func(t *testing.T) {
		userRepository.On("GetByEmail", models.LoginRequest{}).Return(models.User{}, nil).Once()

		result, err := userService.Login(models.LoginRequest{})

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("GetByEmail | Invalid", func(t *testing.T) {
		userRepository.On("GetByEmail", models.LoginRequest{}).Return(models.User{}, errors.New("something went wrong")).Once()

		result, err := userService.Login(models.LoginRequest{})

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestGetUserInfo(t *testing.T) {
	t.Run("GetUserInfo | Valid", func(t *testing.T) {
		userRepository.On("GetUserInfo", "1").Return(models.User{}, nil).Once()

		result, err := userService.GetUserInfo("1")

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("GetUserInfo | Invalid", func(t *testing.T) {
		userRepository.On("GetUserInfo", "0").Return(models.User{}, errors.New("something went wrong")).Once()

		result, err := userService.GetUserInfo("0")

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}
