package controllers

import (
	"time"

	"github.com/labstack/echo/v4"
)

// type Response[T any] struct {
// 	Status  string `json:"status"`
// 	Message string `json:"message"`
// 	Data    T      `json:"data"`
// }

// func NewResponse[T any](c echo.Context, statusCode int, statusMessage string, message string, data T) error {
// 	return c.JSON(statusCode, Response[T]{
// 		Status:  statusMessage,
// 		Message: message,
// 		Data:    data,
// 	})
// }

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
	Meta        Meta   `json:"meta"`
	AccessToken string `json:"token"`
}

func NewResponse[T any](c echo.Context, statusCode int, statusMessage string, message string, data T) error {
	// Jika `data` adalah string (asumsi token JWT), gunakan TokenResponse
	if token, ok := any(data).(string); ok && token != "" {
		return c.JSON(statusCode, TokenResponse{
			Meta: Meta{
				Status:  statusCode,
				Message: message,
				Cache:   false,
				Time:    time.Now().Unix(),
			},
			AccessToken: token,
		})
	}

	// Default response jika bukan token
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
