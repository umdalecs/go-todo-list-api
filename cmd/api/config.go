package main

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var Envs = initConfig()

type Configuration struct {
	JwtSecret string
}

func initConfig() *Configuration {
	godotenv.Load()
	return &Configuration{
		JwtSecret: loadEnv("JWT_SECRET"),
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
