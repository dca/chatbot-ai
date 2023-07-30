package config

import (
	"log"

	"github.com/spf13/viper"
)

// Config is a struct that holds the configuration for the application
type Config struct {
	ChatbotPort string `mapstructure:"CHATBOT_PORT"`

	OPENAI_API_KEY string `mapstructure:"OPENAI_API_KEY"`

	LINE_APP_ID     string `mapstructure:"LINE_CHANNEL_ID"`
	LINE_APP_SECERT string `mapstructure:"LINE_CHANNEL_TOKEN"`
}

// LoadConfig creates a new Config struct
func LoadConfig(path string) (*Config, error) {
	viper.AddConfigPath(path)

	// config file name is: config.env
	viper.SetConfigName("config")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Println(err)
		// allow start without config file
		// return nil, err
	}

	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

// NewTestConfig creates a new Config struct for testing
func NewTestConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.AddConfigPath("../config")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
