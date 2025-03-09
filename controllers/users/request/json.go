package request

import (
	"github.com/amdrx480/go-lms/businesses/users"

	"github.com/go-playground/validator/v10"
)

type UserRegister struct {
	FullName string `json:"fullname" validate:"required"`
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type UserLogin struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (req *UserRegister) ToDomain() *users.Domain {
	return &users.Domain{
		FullName: req.FullName,
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}
}

func (req *UserLogin) ToDomain() *users.Domain {
	return &users.Domain{
		Email:    req.Email,
		Password: req.Password,
	}
}

func (req *UserRegister) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)

	return err
}

func (req *UserLogin) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)

	return err
}
