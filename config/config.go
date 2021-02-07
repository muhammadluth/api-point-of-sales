package config

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"api-point-of-sales/model"

	"github.com/joho/godotenv"
)

func LoadConfig() model.Properties {
	timestart := time.Now()
	fmt.Println("Starting Load Config " + timestart.Format("2006-01-02 15:04:05"))
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(err, "init get config")
		time.Sleep(5 * time.Second)
		os.Exit(1)
	}
	parsePoolSize, err := strconv.Atoi(os.Getenv("POOL_CONNECTION"))
	if err != nil {
		fmt.Println(err, "init get config")
		time.Sleep(5 * time.Second)
		os.Exit(1)
	}
	properties := model.Properties{
		DBHost:             os.Getenv("DB_HOST"),
		DBPort:             os.Getenv("DB_PORT"),
		DBName:             os.Getenv("DB_NAME"),
		DBUser:             os.Getenv("DB_USER"),
		DBPassword:         os.Getenv("DB_PASSWORD"),
		LogPath:            os.Getenv("LOG_PATH"),
		PoolSize:           parsePoolSize,
		PrivateKey:         os.Getenv("PRIVATE_KEY"),
		PublicKey:          os.Getenv("PUBLIC_KEY"),
		ExpireAccessToken:  os.Getenv("EXPIRE_ACCESS_TOKEN"),
		ExpireRefreshToken: os.Getenv("EXPIRE_REFRESH_TOKEN"),
	}
	timefinish := time.Now()
	fmt.Println("Finish Load Config " + timefinish.Format("2006-01-02 15:04:05"))
	return properties
}
