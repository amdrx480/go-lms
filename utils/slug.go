package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"regexp"
	"strings"
	"time"
)

// GenerateSlug membuat slug unik yang bisa berubah jika nama berubah
func GenerateSlug(title string) string {
	// Ubah ke huruf kecil
	slug := strings.ToLower(title)

	// Hapus karakter yang tidak diinginkan (hanya huruf, angka, dan spasi)
	re := regexp.MustCompile(`[^\w\s-]`)
	slug = re.ReplaceAllString(slug, "")

	// Ganti spasi dengan "-"
	slug = strings.ReplaceAll(slug, " ", "-")

	// Hapus karakter "-" ganda
	slug = regexp.MustCompile(`-+`).ReplaceAllString(slug, "-")

	// Tambahkan hash pendek untuk memastikan slug unik
	hash := shortHash(title)

	// Gabungkan slug dengan hash pendek
	return slug + "-" + hash
}

// shortHash membuat hash pendek dari nama agar slug tetap unik
func shortHash(input string) string {
	hash := sha256.Sum256([]byte(input + time.Now().String()))
	return hex.EncodeToString(hash[:])[:8] // Ambil 8 karakter pertama
}
