package helper

import (
	"fmt"
	atypes "github.com/codingbot24-s/types"
	"github.com/golang-jwt/jwt/v5"
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

// check the user token and return the claims

func VerifyTheToken(t string) (atypes.MyClaims, error) {
	secret := GetConfig()
	token, err := jwt.ParseWithClaims(t, &atypes.MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Validate signing algorithm
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret.JwtSecret), nil
	})

	if err != nil {
		return atypes.MyClaims{}, fmt.Errorf("error parsing token: %v", err)
	}

	claims, ok := token.Claims.(*atypes.MyClaims)
	if !ok || !token.Valid {

	}

	return *claims, nil
}
