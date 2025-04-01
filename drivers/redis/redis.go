package redis_driver

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/redis/go-redis/v9"
)

type RedisConfig struct {
	REDIS_HOST     string
	REDIS_PORT     string
	REDIS_PASSWORD string
	REDIS_DB       string
	REDIS_TIMEOUT  string
}

func (config *RedisConfig) InitRedis() *redis.Client {
	addr := fmt.Sprintf("%s:%s", config.REDIS_HOST, config.REDIS_PORT)

	redisDB, err := strconv.Atoi(config.REDIS_DB)
	if err != nil {
		log.Println("Gagal mengonversi SMTP_PORT ke int, menggunakan default (587):", err)
	}

	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: config.REDIS_PASSWORD,
		DB:       redisDB,
	})

	// Cek koneksi Redis
	ctx := context.Background()
	_, err = client.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("error when creating a connection to Redis: %v", err)
	}

	log.Println("Connected to Redis")
	return client
}

func CloseRedis(client *redis.Client) error {
	err := client.Close()
	if err != nil {
		log.Printf("error saat menutup koneksi Redis: %v", err)
		return err
	}

	log.Println("Redis connection closed")
	return nil
}
