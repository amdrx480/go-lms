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
 ┃ ┗ 📂users
 ┃ ┃ ┣ 📂mocks
 ┃ ┃ ┃ ┣ 📜Repository.go
 ┃ ┃ ┃ ┗ 📜Usecase.go
 ┃ ┃ ┣ 📜domain.go
 ┃ ┃ ┣ 📜usecase.go
 ┃ ┃ ┗ 📜usecase_test.go
 ┣ 📂controllers
 ┃ ┣ 📂users
 ┃ ┃ ┣ 📂request
 ┃ ┃ ┃ ┗ 📜json.go
 ┃ ┃ ┣ 📂response
 ┃ ┃ ┃ ┗ 📜json.go
 ┃ ┃ ┗ 📜http.go
 ┃ ┗ 📜base_response.go
 ┣ 📂drivers
 ┃ ┣ 📂mysql
 ┃ ┃ ┣ 📂users
 ┃ ┃ ┃ ┣ 📜mysql.go
 ┃ ┃ ┃ ┗ 📜record.go
 ┃ ┃ ┗ 📜mysql.go
 ┃ ┗ 📜domain_factory.go
 ┣ 📂mariadb
 ┃ ┗ 📜schema.sql
 ┣ 📂tmp
 ┃ ┗ 📜main
 ┣ 📂utils
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
git clone https://github.com/amdrx480/go-clean-architecture-hexagonal.git
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

