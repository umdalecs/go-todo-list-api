package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var Envs = initConfig()

type Configuration struct {
	JwtSecret []byte

	DbAddr string
	DbPort uint16
	DbUser string
	DbPass string
	DbName string
}

func initConfig() *Configuration {
	godotenv.Load()
	return &Configuration{
		JwtSecret: []byte(loadEnv("JWT_SECRET")),

		DbAddr: loadEnv("PG_DB_ADDR"),
		DbPort: uint16(loadIntEnv("PG_DB_PORT")),
		DbUser: loadEnv("PG_DB_USER"),
		DbPass: loadEnv("PG_DB_PASS"),
		DbName: loadEnv("PG_DB_NAME"),
	}
}

func loadEnv(name string) string {
	value, ok := os.LookupEnv(name)
	if !ok {
		log.Fatalf("%s missing", name)
	}

	return value
}

func loadIntEnv(name string) int {
	value, err := strconv.Atoi(loadEnv(name))
	if err != nil {
		log.Fatalf("%s must be an integer", name)
	}
	return value
}
