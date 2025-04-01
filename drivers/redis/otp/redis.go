package otp

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/amdrx480/go-lms/businesses/otp"
	"github.com/redis/go-redis/v9"
)

type OTPRepository struct {
	redisClient *redis.Client
}

func NewOTPRepository(redisClient *redis.Client) otp.Repository {
	return &OTPRepository{
		redisClient: redisClient,
	}
}

func (repo *OTPRepository) SaveOTP(ctx context.Context, otpReq *otp.Domain) (otp.Domain, error) {
	otpRecord := FromDomain(otpReq)

	otpJSON, err := json.Marshal(otpRecord)
	if err != nil {
		return otp.Domain{}, fmt.Errorf("failed to marshal OTP data: %w", err)
	}

	key := fmt.Sprintf("otp:%s", otpReq.Email)
	err = repo.redisClient.Set(ctx, key, otpJSON, time.Until(otpReq.ExpiresAt)).Err()
	if err != nil {
		return otp.Domain{}, fmt.Errorf("failed to save OTP to Redis: %w", err)
	}

	return *otpReq, nil
}

func (repo *OTPRepository) GetOTP(ctx context.Context, otpReq *otp.Domain) (otp.Domain, error) {
	key := fmt.Sprintf("otp:%s", otpReq.Email)
	otpJSON, err := repo.redisClient.Get(ctx, key).Result()
	if err != nil {
		return otp.Domain{}, fmt.Errorf("failed to get OTP from Redis: %w", err)
	}

	var otpRecord OTP
	if err := json.Unmarshal([]byte(otpJSON), &otpRecord); err != nil {
		return otp.Domain{}, fmt.Errorf("failed to unmarshal OTP data: %w", err)
	}

	return otpRecord.ToDomain(), nil
}
