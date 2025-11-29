package helper

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
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

// websocket helper

func HandleConnection(c *websocket.Conn,ch chan string) {
	defer func() {
		c.Close()
		ch <- "done"
	}()
	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		fmt.Printf("recv: %s\n", message)
	}	
}
