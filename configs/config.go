package configs

import (
	"fmt"
	"os"
	"path"

	"github.com/joho/godotenv"
)

var config Config

type Config struct {
	Server   Server
	Database Database
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

func init() {
	work_dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	envPath := path.Join(work_dir, ".env")
	err = godotenv.Load(envPath)
	if err != nil {
		fmt.Println("Error loading .env file")
	}
}
func GetConfig() *Config {
	return &config
}