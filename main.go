package main

import (
	"fmt"
	"github.com/spf13/viper"
)

func init() {
	setUpConfig()
}

func main() {
	fmt.Println("Earn binance funding fee!")
}

func setUpConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config/")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}
