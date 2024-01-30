package configs

import (
	"bytes"
	_ "embed"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

//go:embed config.json
var configFile []byte
var config Config

type Config struct {
	Server   Server
	Database Database
	JWT JWT
}

type Server struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

type Database struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type JWT struct {
	SecretKey string `json:"secret_key"`
}

func init() {
	work_dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// fmt.Println("work_dir", work_dir)
	configSettings := viper.New()

	envPath := path.Join(work_dir, ".env")
	err = godotenv.Load(envPath)
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	// databaseURL := os.Getenv("DATABASE_URL")
	// fmt.Println("ini database url",databaseURL)

	// Environment variables
	configSettings.AutomaticEnv()
	configSettings.SetEnvPrefix("APP")
	configSettings.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))

	//Configurations file
	configSettings.SetConfigType("json")

	err = configSettings.ReadConfig(bytes.NewBuffer(configFile))
	if err != nil {
		fmt.Println("Error reading config file")
	}

	err = configSettings.Unmarshal(&config)
	if err != nil {
		fmt.Println("Error unmarshalling config file")
	}
}

func GetConfig() *Config {
	return &config
}