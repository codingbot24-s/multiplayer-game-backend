package helper

import (
	
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Port      string
	JwtSecret string
}

func NewConfig() {
	// we can create a new config user viper
	viper.SetConfigName("env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading env file, %s", err)
	}
	_ = &Config{
		Port:      viper.GetString("PORT"),
		JwtSecret: viper.GetString("JWT_SECRET"),
	}
}

func GetConfig() *Config {
	return &Config{}
}
