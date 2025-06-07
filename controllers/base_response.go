package controllers

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Meta struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Cache   bool   `json:"cache"`
	Time    int64  `json:"time"`
}

type Response[T any] struct {
	Meta Meta `json:"meta"`
	Data T    `json:"data,omitempty"`
}

type TokenResponse struct {
	Meta         Meta   `json:"meta"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type AuthTokenData struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func NewResponse[T any](c echo.Context, statusCode int, statusMessage string, message string, data T) error {
	switch v := any(data).(type) {
	case AuthTokenData:
		return c.JSON(statusCode, TokenResponse{
			Meta: Meta{
				Status:  statusCode,
				Message: message,
				Cache:   false,
				Time:    time.Now().Unix(),
			},
			AccessToken:  v.AccessToken,
			RefreshToken: v.RefreshToken,
		})
	case *AuthTokenData:
		return c.JSON(statusCode, TokenResponse{
			Meta: Meta{
				Status:  statusCode,
				Message: message,
				Cache:   false,
				Time:    time.Now().Unix(),
			},
			AccessToken:  v.AccessToken,
			RefreshToken: v.RefreshToken,
		})
	default:
		return c.JSON(statusCode, Response[T]{
			Meta: Meta{
				Status:  statusCode,
				Message: message,
				Cache:   false,
				Time:    time.Now().Unix(),
			},
			Data: data,
		})
	}
}

// func NewResponse[T any](c echo.Context, statusCode int, statusMessage string, message string, data T) error {
// 	if tokenData, ok := any(data).(AuthTokenData); ok {
// 		return c.JSON(statusCode, TokenResponse{
// 			Meta: Meta{
// 				Status:  statusCode,
// 				Message: message,
// 				Cache:   false,
// 				Time:    time.Now().Unix(),
// 			},
// 			AccessToken:  tokenData.AccessToken,
// 			RefreshToken: tokenData.RefreshToken,
// 		})
// 	}

// 	// Default response jika bukan token
// 	return c.JSON(statusCode, Response[T]{
// 		Meta: Meta{
// 			Status:  statusCode,
// 			Message: message,
// 			Cache:   false,
// 			Time:    time.Now().Unix(),
// 		},
// 		Data: data,
// 	})
// }

// func NewResponse[T any](c echo.Context, statusCode int, statusMessage string, message string, data T) error {
// 	return c.JSON(statusCode, Response[T]{
// 		Meta: Meta{
// 			Status:  statusCode,
// 			Message: message,
// 			Cache:   false,             // You can decide the appropriate value for 'Cache'
// 			Time:    time.Now().Unix(), // Current timestamp for 'Time'
// 		},
// 		Data: data,
// 	})
// }
