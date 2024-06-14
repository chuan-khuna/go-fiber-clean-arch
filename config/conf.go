package config

import (
	"os"

	"github.com/joho/godotenv"
)

var Port = getPort()

var DB = getDB()
var DBHost = getDBHost()
var DBPort = getDBPort()
var DBUser = getDBUser()
var DBPassword = getDBPassword()
var SecretKey = getSercretKey()

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

}

func getPort() string {
	loadEnv()
	port := os.Getenv("FIBER_PORT")
	if port == "" {
		// Default port
		return "3100"
	}
	return port
}

func getSercretKey() string {
	loadEnv()
	secretKey := os.Getenv("SECRET_KEY")
	if secretKey == "" {
		// Default
		return "Pj9t31C13m0y9n934VyoyBRKZY48Lgh2"
	}
	return secretKey
}

func getDB() string {
	loadEnv()
	db := os.Getenv("FIBER_POSTGRES_DB")
	return db
}

func getDBUser() string {
	loadEnv()
	user := os.Getenv("FIBER_POSTGRES_USER")
	return user
}

func getDBPassword() string {
	loadEnv()
	password := os.Getenv("FIBER_POSTGRES_PASSWORD")
	return password
}

func getDBHost() string {
	loadEnv()
	host := os.Getenv("FIBER_POSTGRES_HOST")
	return host
}

func getDBPort() string {
	loadEnv()
	port := os.Getenv("FIBER_POSTGRES_PORT")
	// convert port to string
	if port == "" {
		// Default port
		return "5432"
	}

	return port
}
