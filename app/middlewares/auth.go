package middlewares

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/amdrx480/angsana-boga/utils"
	"github.com/golang-jwt/jwt/v4"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type JWTCustomClaims struct {
	ID   int        `json:"id"`
	Role utils.Role `json:"role"`
	jwt.RegisteredClaims
}

type JWTConfig struct {
	SecretKey            string
	AccessTokenDuration  time.Duration // e.g. 15 * time.Minute
	RefreshTokenDuration time.Duration // e.g. 7 * 24 * time.Hour
	ExpiresDuration      int
}

type contextKey string

const userContextKey = contextKey("user")

func (jwtConfig *JWTConfig) Init() echojwt.Config {
	return echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(JWTCustomClaims)
		},
		SigningKey: []byte(jwtConfig.SecretKey),
	}
}

func (jwtConfig *JWTConfig) GenerateAccessToken(userID int, role utils.Role) (string, error) {
	claims := &JWTCustomClaims{
		ID:   userID,
		Role: role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(jwtConfig.AccessTokenDuration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := rawToken.SignedString([]byte(jwtConfig.SecretKey))

	if err != nil {
		return "", err
	}

	return token, nil
}

func (jwtConfig *JWTConfig) GenerateRefreshToken(userID int, role utils.Role) (string, error) {
	claims := &JWTCustomClaims{
		ID:   userID,
		Role: role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(jwtConfig.RefreshTokenDuration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := rawToken.SignedString([]byte(jwtConfig.SecretKey))

	if err != nil {
		return "", err
	}

	return token, nil
}

func GetUser(ctx context.Context) (*JWTCustomClaims, error) {
	user, ok := ctx.Value(userContextKey).(*jwt.Token)
	if !ok || user == nil {
		return nil, errors.New("invalid token")
	}

	claims, ok := user.Claims.(*JWTCustomClaims)
	if !ok {
		return nil, errors.New("invalid claims")
	}

	return claims, nil
}

func GetUserID(ctx context.Context) (int, error) {
	claim, err := GetUser(ctx)

	if err != nil {
		return 0, errors.New("invalid token")
	}

	return claim.ID, nil
}

func VerifyToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)

		if user == nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"message": "invalid token",
			})
		}

		ctx := context.WithValue(c.Request().Context(), userContextKey, user)
		c.SetRequest(c.Request().WithContext(ctx))

		userData, err := GetUser(ctx)
		if userData == nil || err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"message": "invalid token",
			})
		}

		return next(c)
	}
}

func (jwtConfig *JWTConfig) VerifyRefreshToken(refreshToken string) (*JWTCustomClaims, error) {
	// Parse token dengan claims custom
	token, err := jwt.ParseWithClaims(refreshToken, &JWTCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Pastikan pakai metode signing HS256
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtConfig.SecretKey), nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to parse refresh token: %w", err)
	}

	// Cek validitas token dan klaim
	if claims, ok := token.Claims.(*JWTCustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, errors.New("invalid refresh token")
	}
}

func VerifyAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user, err := GetUser(c.Request().Context())

		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"message": "invalid token",
			})
		}

		if user.Role != utils.ROLE_ADMIN {
			return c.JSON(http.StatusForbidden, map[string]string{
				"message": "access denied",
			})
		}

		return next(c)
	}
}

func VerifyUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user, err := GetUser(c.Request().Context())

		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"message": "invalid token",
			})
		}

		if user.Role != utils.ROLE_USER {
			return c.JSON(http.StatusForbidden, map[string]string{
				"message": "access denied",
			})
		}

		return next(c)
	}
}
