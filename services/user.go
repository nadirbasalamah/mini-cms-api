package services

import (
	"mini-cms-api/models"
	"mini-cms-api/repositories"
	"mini-cms-api/utils"
)

type UserService struct {
	Repository repositories.UserRepository
	JwtOptions models.JWTOptions
}

func InitUserService(jwtOptions models.JWTOptions) UserService {
	return UserService{
		Repository: &repositories.UserRepositoryImpl{},
		JwtOptions: jwtOptions,
	}
}

func (us *UserService) Register(registerReq models.RegisterRequest) (models.User, error) {
	return us.Repository.Register(registerReq)
}

func (us *UserService) Login(loginReq models.LoginRequest) (string, error) {
	user, err := us.Repository.GetByEmail(loginReq)

	if err != nil {
		return "", err
	}

	token, err := utils.GenerateJWT(int(user.ID), us.JwtOptions)

	if err != nil {
		return "", err
	}

	return token, nil
}

func (us *UserService) GetUserInfo(id string) (models.User, error) {
	return us.Repository.GetUserInfo(id)
}
