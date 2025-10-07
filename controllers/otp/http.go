package otp

import (
	"log"
	"net/http"

	"github.com/amdrx480/angsana-boga/businesses/otp"
	"github.com/amdrx480/angsana-boga/controllers"
	"github.com/amdrx480/angsana-boga/controllers/otp/request"

	"github.com/labstack/echo/v4"
)

type OTPController struct {
	otpUseCase otp.UseCase
}

func NewOTPController(lessonUC otp.UseCase) *OTPController {
	return &OTPController{
		otpUseCase: lessonUC,
	}
}

func (oc *OTPController) RequestOTP(c echo.Context) error {
	otpReq := request.OTPRequest{}
	ctx := c.Request().Context()

	if err := c.Bind(&otpReq); err != nil {
		log.Println("Error binding request:", err)
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "invalid request", "")
	}

	if err := otpReq.Validate(); err != nil {
		log.Println("Validation error:", err)
		return controllers.NewResponse(c, http.StatusUnprocessableEntity, "failed", err.Error(), "")
	}

	// otpData, err := oc.otpUseCase.RequestOTP(ctx, otpReq.ToDomain())
	_, err := oc.otpUseCase.RequestOTP(ctx, otpReq.ToDomain())
	if err != nil {
		log.Println("Failed to create OTP:", err) // Tambahkan log error
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "failed to create a otp", "")
	}

	return controllers.NewResponse(c, http.StatusCreated, "success", "otp created", "")
}

func (oc *OTPController) LoginWithOTP(c echo.Context) error {
	loginReq := request.LoginWithOTP{}
	ctx := c.Request().Context()

	// Bind request JSON ke struct
	if err := c.Bind(&loginReq); err != nil {
		log.Println("Error binding request:", err)
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "invalid request", "")
	}

	// Validasi input
	if err := loginReq.Validate(); err != nil {
		log.Println("Validation error:", err)
		return controllers.NewResponse(c, http.StatusUnprocessableEntity, "failed", err.Error(), "")
	}

	// Konversi ke domain yang sesuai dengan OTPUseCase
	otpDomain := otp.Domain{
		Email:   loginReq.Email,
		OTPCode: loginReq.OTPCode,
	}

	// Cek OTP
	token, err := oc.otpUseCase.LoginWithOTP(ctx, &otpDomain)

	if err != nil {
		log.Println("❌ Failed to verify OTP for email:", loginReq.Email, "Error:", err)

		if err.Error() == "OTP tidak valid" || err.Error() == "OTP tidak ditemukan atau sudah kedaluwarsa" {
			return controllers.NewResponse(c, http.StatusUnauthorized, "failed", err.Error(), "")
		}

		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "failed to create a otp", "")
	}

	log.Println("✅ OTP verified successfully for email:", loginReq.Email)

	return controllers.NewResponse(c, http.StatusOK, "success", "OTP verified", token)
}
