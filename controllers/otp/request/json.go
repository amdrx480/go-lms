package request

import (
	"github.com/amdrx480/angsana-boga/businesses/otp"
	"github.com/amdrx480/angsana-boga/businesses/users"
	"github.com/go-playground/validator/v10"
)

type OTPRequest struct {
	Email string `json:"email" validate:"required"`
}

type LoginWithOTP struct {
	Email   string `json:"email" validate:"required"`
	OTPCode string `json:"otp_code" validate:"required"`
}

func (req *OTPRequest) ToDomain() *users.Domain {
	return &users.Domain{
		Email: req.Email,
	}
}

func (req *LoginWithOTP) ToOTPDomain() *otp.Domain {
	return &otp.Domain{
		Email:   req.Email,
		OTPCode: req.OTPCode,
	}
}
func (req *OTPRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)

	return err
}

func (req *LoginWithOTP) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)

	return err
}
