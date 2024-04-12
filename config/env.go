package config

import "github.com/joho/godotenv"

func InitEnv() {
	if err := godotenv.Load(".env"); err != nil {
		panic("No se encontró archivo '.env'")
	}
}
