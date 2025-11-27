package helper

import (
	"fmt"
	"log"
	"net/http"

	atypes "github.com/codingbot24-s/types"
	"github.com/fasthttp/websocket"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
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

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func ConnectToWS(zone string, r *http.Request, w http.ResponseWriter) error {
	// here we need to connect to websocket
	_, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return fmt.Errorf("error upgrading to websocket")
	}

	// is this the right thing to do here ?
	return nil
}
