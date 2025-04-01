package utils

import (
	"crypto/rand"
	"errors"
	"fmt"
	"log"
	"math/big"
	"strconv"
	"time"

	mail "github.com/xhit/go-simple-mail/v2"
)

// GenerateRandomOTP menghasilkan OTP acak 6 digit
func GenerateRandomOTP() string {
	n, err := rand.Int(rand.Reader, big.NewInt(1000000))
	if err != nil {
		return "000000" // Fallback jika terjadi error
	}
	return fmt.Sprintf("%06d", n.Int64())
}

var smtpClient *mail.SMTPClient

// SMTPConfig menyimpan konfigurasi SMTP
type SMTPConfig struct {
	SMTP_HOST     string
	SMTP_PORT     string
	SMTP_EMAIL    string
	SMTP_PASSWORD string
	SMTP_NAME     string
	SMTP_TIMEOUT  string
}

// **Inisialisasi SMTP Client**
func (config *SMTPConfig) InitSMTP() *mail.SMTPClient {
	if smtpClient != nil {
		log.Println("SMTP sudah terhubung, menggunakan koneksi yang ada")
		return smtpClient
	}

	port, err := strconv.Atoi(config.SMTP_PORT)
	if err != nil {
		log.Println("Kesalahan parsing SMTP_PORT:", err)
		return nil
	}

	server := mail.NewSMTPClient()
	server.Host = config.SMTP_HOST
	server.Port = port
	server.Username = config.SMTP_EMAIL
	server.Password = config.SMTP_PASSWORD
	server.Encryption = mail.EncryptionSTARTTLS
	server.KeepAlive = true // Tetap hidup untuk pengiriman berulang
	server.ConnectTimeout = 10 * time.Second
	server.SendTimeout = 10 * time.Second

	// Hubungkan ke SMTP
	smtpClient, err = server.Connect()
	if err != nil {
		log.Println("Gagal menghubungkan ke SMTP:", err)
		return nil
	}

	log.Println("Terhubung ke SMTP")
	return smtpClient
}

// **Mengirim Email OTP**
func SendOTPEmail(recipientEmail, otpCode string) error {
	if smtpClient == nil {
		return errors.New("SMTP belum diinisialisasi")
	}

	// Buat email
	email := mail.NewMSG()
	email.SetFrom(fmt.Sprintf("%s <%s>", GetConfig("SMTP_NAME"), GetConfig("SMTP_EMAIL")))
	email.AddTo(recipientEmail)
	email.SetSubject("Kode OTP Anda")
	email.SetBody(mail.TextPlain, fmt.Sprintf("Kode OTP Anda adalah: %s", otpCode))

	// Kirim email
	err := email.Send(smtpClient)
	if err != nil {
		log.Println("Gagal mengirim email OTP:", err)
		return err
	}

	log.Println("Email OTP berhasil dikirim ke:", recipientEmail)
	return nil
}

// **Menutup Koneksi SMTP**
func CloseSMTP(client *mail.SMTPClient) error {
	if client == nil {
		return errors.New("SMTP client tidak tersedia")
	}

	client.Close()
	log.Println("SMTP connection closed")
	smtpClient = nil
	return nil
}
