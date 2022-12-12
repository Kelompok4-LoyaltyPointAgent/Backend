package config

import (
	"log"
	"os"
	"strconv"
)

type RedisConfig struct {
	Addr     string
	Port     string
	Username string
	Password string
	DB       int
}

func LoadRedisConfig() RedisConfig {
	db, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	if err != nil {
		log.Print("Invalid Redis DB. Default value will be used!")
	}

	return RedisConfig{
		Addr:     os.Getenv("REDIS_ADDR"),
		Port:     os.Getenv("REDIS_PORT"),
		Username: os.Getenv("REDIS_USERNAME"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       db,
	}
}
