package users

import (
	"context"
	"errors"
	"fmt"

	"github.com/amdrx480/go-lms/app/middlewares"
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

func (usecase *userUseCase) Login(ctx context.Context, userDomain *Domain) (string, string, error) {
	user, err := usecase.userRepository.GetByEmail(ctx, userDomain)

	if err != nil {
		return "", "", errors.New("user not found")
	}

	// Generate Access Token
	accessToken, err := usecase.jwtConfig.GenerateAccessToken(int(user.ID), user.Role)

	if err != nil {
		return "", "", errors.New("gagal membuat access token")
	}

	// Generate Refresh Token
	refreshToken, err := usecase.jwtConfig.GenerateRefreshToken(user.ID, user.Role)
	if err != nil {
		return "", "", errors.New("gagal membuat refresh token")
	}

	return accessToken, refreshToken, nil
}

func (usecase *userUseCase) RefreshAccessToken(ctx context.Context, refreshToken string) (string, error) {
	claims, err := usecase.jwtConfig.VerifyRefreshToken(refreshToken)
	if err != nil {
		return "", fmt.Errorf("invalid refresh token: %w", err)
	}

	newAccessToken, err := usecase.jwtConfig.GenerateAccessToken(claims.ID, claims.Role)
	if err != nil {
		return "", fmt.Errorf("failed to generate new access token: %w", err)
	}

	return newAccessToken, nil
}

func (usecase *userUseCase) GetUserProfile(ctx context.Context) (Domain, error) {
	return usecase.userRepository.GetUserProfile(ctx)
}
