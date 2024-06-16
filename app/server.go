package app

import (
	"flag"
	"log"
	"os"

	"github.com/Gswaraw37/toko-online/app/controllers"
	"github.com/joho/godotenv"
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func Run() {
	server := controllers.Server{}
	appConfig := controllers.AppConfig{}
	dbConfig := controllers.DBConfig{}

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error .env file")
	}

	appConfig.AppName = getEnv("APP_NAME", "Toko Online")
	appConfig.AppEnv = getEnv("APP_ENV", "development")
	appConfig.AppPort = getEnv("APP_PORT", "9000")
	appConfig.AppURL = getEnv("APP_URL", "http://localhost:9000")

	dbConfig.DBDriver = getEnv("DB_DRIVER", "postgres")
	dbConfig.DBHost = getEnv("DB_HOST", "localhost")
	dbConfig.DBUser = getEnv("DB_USER", "postgres")
	dbConfig.DBPassword = getEnv("DB_PASSWORD", "postgres")
	dbConfig.DBName = getEnv("DB_NAME", "dbname")
	dbConfig.DBPort = getEnv("DB_PORT", "5432")

	flag.Parse()
	arg := flag.Arg(0)

	if arg != "" {
		server.InitCommands(appConfig, dbConfig)
	} else {
		server.Initialize(appConfig, dbConfig)
		server.Run(":" + appConfig.AppPort)
	}
}
