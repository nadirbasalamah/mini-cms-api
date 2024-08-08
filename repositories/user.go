package repositories

import (
	"mini-cms-api/database"
	"mini-cms-api/models"

	"golang.org/x/crypto/bcrypt"
)

type UserRepositoryImpl struct{}

func InitUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (ur *UserRepositoryImpl) Register(registerReq models.RegisterRequest) (models.User, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(registerReq.Password), bcrypt.DefaultCost)

	if err != nil {
		return models.User{}, err
	}

	var user models.User = models.User{
		Email:    registerReq.Email,
		Password: string(password),
	}

	result := database.DB.Create(&user)

	if err := result.Error; err != nil {
		return models.User{}, err
	}

	err = result.Last(&user).Error

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (ur *UserRepositoryImpl) GetByEmail(loginReq models.LoginRequest) (models.User, error) {
	var user models.User

	err := database.DB.First(&user, "email = ?", loginReq.Email).Error

	if err != nil {
		return models.User{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginReq.Password))

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (ur *UserRepositoryImpl) GetUserInfo(id string) (models.User, error) {
	var user models.User

	err := database.DB.First(&user, "id = ?", id).Error

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
