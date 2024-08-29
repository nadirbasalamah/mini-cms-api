package users

import "mini-cms-api/app/middlewares"

type userUseCase struct {
	userRepository Repository
	jwtAuth        *middlewares.JWTConfig
}

func NewUserUseCase(repository Repository, jwtAuth *middlewares.JWTConfig) UseCase {
	return &userUseCase{
		userRepository: repository,
		jwtAuth:        jwtAuth,
	}
}

func (usecase *userUseCase) Register(userReq *Domain) (Domain, error) {
	return usecase.userRepository.Register(userReq)
}

func (usecase *userUseCase) Login(userReq *Domain) (string, error) {
	user, err := usecase.userRepository.GetUserByEmail(userReq)

	if err != nil {
		return "", err
	}

	token, err := usecase.jwtAuth.GenerateToken(int(user.ID))

	if err != nil {
		return "", err
	}

	return token, nil
}

func (usecase *userUseCase) GetUserInfo(id string) (Domain, error) {
	return usecase.userRepository.GetUserInfo(id)
}
