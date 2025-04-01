package response

import (
	"time"

	"github.com/amdrx480/go-lms/businesses/otp"
)

type OTP struct {
	Email     string    `json:"email"`
	OTPCode   string    `json:"otp_code"`
	ExpiresAt time.Time `json:"expires_at"`
}

func FromDomain(domain otp.Domain) *OTP {
	return &OTP{
		Email:     domain.Email,
		OTPCode:   domain.OTPCode,
		ExpiresAt: domain.ExpiresAt,
	}
}
