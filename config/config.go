package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	App struct {
		Name string `yaml:"name"`
		Port string    `yaml:"port"`
	} `yaml:"app"`
	Database struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Name     string `yaml:"name"`
		MaxIdleConns	int    `yaml:"max_idle_conns"`
		MaxOpenConns	int    `yaml:"max_open_conns"`
	} `yaml:"database"`
	JWT struct {
		Secret         string `yaml:"secret"`
		ExpireDuration int    `yaml:"expire_duration"`
	} `yaml:"jwt"`
}

var AppConfig Config

func InitConfig() {
	// Configuration initialization logic (e.g., reading from a YAML file)
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	AppConfig = Config{}

	if err := viper.Unmarshal(&AppConfig); err != nil {
		log.Fatalf("Unable to decode into struct: %v", err)
	}

	initDB()
}