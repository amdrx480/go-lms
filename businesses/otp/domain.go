package otp

import (
	"context"
	"time"

	"github.com/amdrx480/angsana-boga/businesses/users"
)

type Domain struct {
	Email     string
	OTPCode   string
	ExpiresAt time.Time
}

type UseCase interface {
	RequestOTP(ctx context.Context, userReq *users.Domain) (Domain, error)
	LoginWithOTP(ctx context.Context, otpReq *Domain) (string, error)
}

type Repository interface {
	SaveOTP(ctx context.Context, otpReq *Domain) (Domain, error)
	GetOTP(ctx context.Context, otpReq *Domain) (Domain, error)
}
