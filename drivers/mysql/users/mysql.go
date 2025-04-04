package users

import (
	"context"
	"errors"

	"github.com/amdrx480/go-lms/app/middlewares"
	"github.com/amdrx480/go-lms/businesses/users"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type userRepository struct {
	conn *gorm.DB
}

func NewMySQLRepository(conn *gorm.DB) users.Repository {
	return &userRepository{
		conn: conn,
	}
}

func (ur *userRepository) Register(ctx context.Context, userDomain *users.Domain) (users.Domain, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(userDomain.Password), bcrypt.DefaultCost)

	if err != nil {
		return users.Domain{}, err
	}

	record := FromDomain(userDomain)

	record.Password = string(password)

	result := ur.conn.WithContext(ctx).Create(&record)

	if err := result.Error; err != nil {
		return users.Domain{}, err
	}

	err = result.Last(&record).Error

	if err != nil {
		return users.Domain{}, err
	}

	return record.ToDomain(), nil
}

func (ur *userRepository) GetByEmail(ctx context.Context, userDomain *users.Domain) (users.Domain, error) {
	var user User

	err := ur.conn.WithContext(ctx).First(&user, "email = ?", userDomain.Email).Error

	if err != nil {
		return users.Domain{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userDomain.Password))

	if err != nil {
		return users.Domain{}, err
	}

	return user.ToDomain(), nil
}

func (ur *userRepository) FindByEmail(ctx context.Context, email string) (users.Domain, error) {
	var user User

	err := ur.conn.WithContext(ctx).First(&user, "email = ?", email).Error

	if err != nil {
		return users.Domain{}, err
	}

	return user.ToDomain(), nil
}

func (ur *userRepository) GetUserProfile(ctx context.Context) (users.Domain, error) {
	id, err := middlewares.GetUserID(ctx)

	if err != nil {
		return users.Domain{}, errors.New("invalid token")
	}

	var user User

	err = ur.conn.WithContext(ctx).First(&user, "id = ?", id).Error

	if err != nil {
		return users.Domain{}, err
	}

	return user.ToDomain(), nil
}
