package users

import (
	"context"

	"github.com/amdrx480/go-clean-architecture-hexagonal/app/middlewares"
)

type userUseCase struct {
	userRepository Repository
	jwtConfig      *middlewares.JWTConfig
}

func NewUserUseCase(repository Repository, jwtConfig *middlewares.JWTConfig) UseCase {
	return &userUseCase{
		userRepository: repository,
		jwtConfig:      jwtConfig,
	}
}

func (usecase *userUseCase) Register(ctx context.Context, userDomain *Domain) (Domain, error) {
	return usecase.userRepository.Register(ctx, userDomain)
}

func (usecase *userUseCase) Login(ctx context.Context, userDomain *Domain) (string, error) {
	user, err := usecase.userRepository.GetByEmail(ctx, userDomain)

	if err != nil {
		return "", err
	}

	token, err := usecase.jwtConfig.GenerateToken(int(user.ID), user.Role)

	if err != nil {
		return "", err
	}

	return token, nil
}

func (usecase *userUseCase) GetUserProfile(ctx context.Context) (Domain, error) {
	return usecase.userRepository.GetUserProfile(ctx)
}
