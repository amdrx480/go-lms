package otp

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/amdrx480/go-lms/app/middlewares"
	"github.com/amdrx480/go-lms/businesses/users"
	"github.com/amdrx480/go-lms/utils"
)

type OTPUseCase struct {
	userRepository users.Repository
	otpRepository  Repository
	jwtConfig      *middlewares.JWTConfig
}

func NewOTPUseCase(userRepository users.Repository, otpRepository Repository, jwtConfig *middlewares.JWTConfig) UseCase {
	return &OTPUseCase{
		userRepository: userRepository,
		otpRepository:  otpRepository,
		jwtConfig:      jwtConfig,
	}
}

func (usecase *OTPUseCase) RequestOTP(ctx context.Context, userReq *users.Domain) (Domain, error) {
	log.Println("Mencari user dengan email:", userReq.Email) // Tambahkan log ini

	user, err := usecase.userRepository.FindByEmail(ctx, userReq.Email)
	if err != nil {
		log.Println("User tidak ditemukan untuk email:", userReq.Email) // Tambahkan log ini
		return Domain{}, errors.New("user tidak ditemukan")
	}

	otpCode := utils.GenerateRandomOTP()

	otpData := Domain{
		Email:     user.Email,
		OTPCode:   otpCode,
		ExpiresAt: time.Now().Add(5 * time.Minute),
	}

	_, err = usecase.otpRepository.SaveOTP(ctx, &otpData)
	if err != nil {
		log.Println("Gagal menyimpan OTP untuk email:", otpData.Email) // Tambahkan log ini
		return Domain{}, errors.New("gagal menyimpan OTP")
	}

	log.Println("OTP berhasil dibuat untuk email:", otpData.Email) // Tambahkan log ini

	// Kirim email OTP
	err = utils.SendOTPEmail(user.Email, otpCode)
	if err != nil {
		log.Println("❌ Gagal mengirim OTP ke email:", err)
		return Domain{}, errors.New("gagal mengirim OTP email")
	}

	log.Println("✅ OTP berhasil dikirim ke:", user.Email)

	return otpData, nil
}

func (usecase *OTPUseCase) LoginWithOTP(ctx context.Context, otpReq *Domain) (string, error) {
	log.Println("🔍 Mencari user dengan email:", otpReq.Email)

	// Cari user berdasarkan email
	user, err := usecase.userRepository.FindByEmail(ctx, otpReq.Email)
	if err != nil {
		log.Println("User tidak ditemukan untuk email:", otpReq.Email)
		return "", errors.New("user tidak ditemukan")
	}

	// Ambil OTP dari Redis
	otpFromRedis, err := usecase.otpRepository.GetOTP(ctx, otpReq) // Langsung gunakan otpReq
	if err != nil {
		log.Println("OTP tidak ditemukan atau telah kedaluwarsa untuk email:", otpReq.Email)
		return "", errors.New("OTP tidak ditemukan atau sudah kedaluwarsa")
	}

	// Periksa apakah OTP cocok
	if otpFromRedis.OTPCode != otpReq.OTPCode {
		log.Println("OTP tidak valid untuk email:", otpReq.Email)
		return "", errors.New("OTP tidak valid")
	}

	log.Println("OTP valid untuk email:", otpReq.Email)

	// Buat token JWT
	token, err := usecase.jwtConfig.GenerateToken(int(user.ID), user.Role)
	if err != nil {
		log.Println("Gagal membuat token JWT untuk email:", otpReq.Email)
		return "", err
	}

	log.Println("✅ Token JWT berhasil dibuat untuk email:", otpReq.Email)

	return token, nil
}
