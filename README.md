# Go Clean Architecture Hexagonal

REST API aplikasi menggunakan bahasa Go dengan Echo Framework.

## 📌 Deskripsi

Aplikasi ini menggunakan arsitektur Clean Architecture dengan pendekatan Hexagonal untuk memastikan kode tetap terstruktur, terpisah, dan mudah diuji. Proyek ini menggunakan GORM sebagai ORM untuk integrasi dengan database MySQL.

## 🛠️ Teknologi yang Digunakan

- Golang
- Echo Framework
- GORM
- MySQL
- JWT Authentication
- Docker
- Viper (Konfigurasi)
- Validator (Validasi Input)
- Testify (Unit Testing)

## 📂 Struktur Folder

```bash
 ┣ 📂app
 ┃ ┣ 📂middlewares
 ┃ ┃ ┣ 📜auth.go
 ┃ ┃ ┗ 📜logger.go
 ┃ ┗ 📂routes
 ┃ ┃ ┗ 📜routes.go
 ┣ 📂businesses
 ┃ ┣ 📂categories
 ┃ ┃ ┣ 📜domain.go
 ┃ ┃ ┗ 📜usecase.go
 ┃ ┣ 📂chapters
 ┃ ┃ ┣ 📜domain.go
 ┃ ┃ ┗ 📜usecase.go
 ┃ ┣ 📂courses
 ┃ ┃ ┣ 📜domain.go
 ┃ ┃ ┗ 📜usecase.go
 ┃ ┣ 📂documents
 ┃ ┃ ┣ 📜domain.go
 ┃ ┃ ┗ 📜usecase.go
 ┃ ┣ 📂enrollments
 ┃ ┃ ┣ 📜domain.go
 ┃ ┃ ┗ 📜usecase.go
 ┃ ┣ 📂lessons
 ┃ ┃ ┣ 📜domain.go
 ┃ ┃ ┗ 📜usecase.go
 ┃ ┣ 📂modules
 ┃ ┃ ┣ 📜domain.go
 ┃ ┃ ┗ 📜usecase.go
 ┃ ┣ 📂otp
 ┃ ┃ ┣ 📜domain.go
 ┃ ┃ ┗ 📜usecase.go
 ┃ ┗ 📂users
 ┃ ┃ ┣ 📂mocks
 ┃ ┃ ┃ ┣ 📜Repository.go
 ┃ ┃ ┃ ┗ 📜Usecase.go
 ┃ ┃ ┣ 📜domain.go
 ┃ ┃ ┣ 📜usecase.go
 ┃ ┃ ┗ 📜usecase_test.go
 ┣ 📂controllers
 ┃ ┣ 📂categories
 ┃ ┃ ┣ 📂request
 ┃ ┃ ┃ ┗ 📜json.go
 ┃ ┃ ┣ 📂response
 ┃ ┃ ┃ ┗ 📜json.go
 ┃ ┃ ┗ 📜http.go
 ┃ ┣ 📂chapters
 ┃ ┃ ┣ 📂request
 ┃ ┃ ┃ ┗ 📜json.go
 ┃ ┃ ┣ 📂response
 ┃ ┃ ┃ ┗ 📜json.go
 ┃ ┃ ┗ 📜http.go
 ┃ ┣ 📂courses
 ┃ ┃ ┣ 📂request
 ┃ ┃ ┃ ┗ 📜json.go
 ┃ ┃ ┣ 📂response
 ┃ ┃ ┃ ┗ 📜json.go
 ┃ ┃ ┗ 📜http.go
 ┃ ┣ 📂documents
 ┃ ┃ ┣ 📂request
 ┃ ┃ ┃ ┗ 📜json.go
 ┃ ┃ ┣ 📂response
 ┃ ┃ ┃ ┗ 📜json.go
 ┃ ┃ ┗ 📜http.go
 ┃ ┣ 📂enrollments
 ┃ ┃ ┣ 📂request
 ┃ ┃ ┃ ┗ 📜json.go
 ┃ ┃ ┣ 📂response
 ┃ ┃ ┃ ┗ 📜json.go
 ┃ ┃ ┗ 📜http.go
 ┃ ┣ 📂lessons
 ┃ ┃ ┣ 📂request
 ┃ ┃ ┃ ┗ 📜json.go
 ┃ ┃ ┣ 📂response
 ┃ ┃ ┃ ┗ 📜json.go
 ┃ ┃ ┗ 📜http.go
 ┃ ┣ 📂modules
 ┃ ┃ ┣ 📂request
 ┃ ┃ ┃ ┗ 📜json.go
 ┃ ┃ ┣ 📂response
 ┃ ┃ ┃ ┗ 📜json.go
 ┃ ┃ ┗ 📜http.go
 ┃ ┣ 📂otp
 ┃ ┃ ┣ 📂request
 ┃ ┃ ┃ ┗ 📜json.go
 ┃ ┃ ┣ 📂response
 ┃ ┃ ┃ ┗ 📜json.go
 ┃ ┃ ┗ 📜http.go
 ┃ ┣ 📂users
 ┃ ┃ ┣ 📂request
 ┃ ┃ ┃ ┗ 📜json.go
 ┃ ┃ ┣ 📂response
 ┃ ┃ ┃ ┗ 📜json.go
 ┃ ┃ ┗ 📜http.go
 ┃ ┗ 📜base_response.go
 ┣ 📂drivers
 ┃ ┣ 📂mysql
 ┃ ┃ ┣ 📂categories
 ┃ ┃ ┃ ┣ 📜mysql.go
 ┃ ┃ ┃ ┗ 📜record.go
 ┃ ┃ ┣ 📂chapters
 ┃ ┃ ┃ ┣ 📜mysql.go
 ┃ ┃ ┃ ┗ 📜record.go
 ┃ ┃ ┣ 📂courses
 ┃ ┃ ┃ ┣ 📜mysql.go
 ┃ ┃ ┃ ┗ 📜record.go
 ┃ ┃ ┣ 📂documents
 ┃ ┃ ┃ ┣ 📜mysql.go
 ┃ ┃ ┃ ┗ 📜record.go
 ┃ ┃ ┣ 📂enrollments
 ┃ ┃ ┃ ┣ 📜mysql.go
 ┃ ┃ ┃ ┗ 📜record.go
 ┃ ┃ ┣ 📂lessons
 ┃ ┃ ┃ ┣ 📜mysql.go
 ┃ ┃ ┃ ┗ 📜record.go
 ┃ ┃ ┣ 📂modules
 ┃ ┃ ┃ ┣ 📜mysql.go
 ┃ ┃ ┃ ┗ 📜record.go
 ┃ ┃ ┣ 📂users
 ┃ ┃ ┃ ┣ 📜mysql.go
 ┃ ┃ ┃ ┗ 📜record.go
 ┃ ┃ ┗ 📜mysql.go
 ┃ ┣ 📂redis
 ┃ ┃ ┣ 📂otp
 ┃ ┃ ┃ ┣ 📜record.go
 ┃ ┃ ┃ ┗ 📜redis.go
 ┃ ┃ ┗ 📜redis.go
 ┃ ┗ 📜domain_factory.go
 ┣ 📂mariadb
 ┃ ┗ 📜schema.sql
 ┣ 📂tmp
 ┃ ┗ 📜main
 ┣ 📂utils
 ┃ ┣ 📜const.go
 ┃ ┣ 📜otp.go
 ┃ ┣ 📜slug.go
 ┃ ┗ 📜utils.go
 ┣ 📜.air.toml
 ┣ 📜.env
 ┣ 📜.env.example
 ┣ 📜.gitignore
 ┣ 📜docker-compose.yml
 ┣ 📜Dockerfile
 ┣ 📜go.mod
 ┣ 📜go.sum
 ┣ 📜main.go
 ┗ 📜README.md
```

## 🔑 Fitur

- Registrasi dan Login Pengguna
- Otentikasi JWT
- Validasi Input
- Penggunaan Middleware
- Unit Testing
- Dockerized Application
- Logging request

## 📌 Instalasi

1. Clone Repository

```bash
git clone https://github.com/amdrx480/go-lms.git
cd go-clean-architecture-hexagonal
```

2. Copy file konfigurasi

```bash
cp .env.example .env
```

3. Edit file `.env` sesuai konfigurasi database.

4. Jalankan aplikasi menggunakan Docker

```bash
docker-compose up --build
```

## 🎯 Testing

Unit testing menggunakan library **Testify**.

```bash
go test ./...
```

