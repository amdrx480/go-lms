package otp

import (
	"time"

	"github.com/amdrx480/angsana-boga/businesses/otp"
)

type OTP struct {
	Email     string    `json:"email"`
	OTPCode   string    `json:"otp_code"`
	ExpiresAt time.Time `json:"expires_at"`
}

func (rec *OTP) ToDomain() otp.Domain {
	return otp.Domain{
		Email:     rec.Email,
		OTPCode:   rec.OTPCode,
		ExpiresAt: rec.ExpiresAt,
	}
}

func FromDomain(domain *otp.Domain) *OTP {
	return &OTP{
		Email:     domain.Email,
		OTPCode:   domain.OTPCode,
		ExpiresAt: domain.ExpiresAt,
	}
}
