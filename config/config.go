package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
)

type Config struct {
	Postgres struct {
		Login    string `yaml:"login"`
		Password string `yaml:"password"`
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		DBName   string `yaml:"db_name"`
		PGDriver string `yaml:"pg_driver"`
		SSLMode  string `yaml:"ssl_mode"`
	} `yaml:"postgres"`
	Mongodb struct {
		URI      string `yaml:"URI"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		DB       string `yaml:"db"`
	} `yaml:"mongodb"`
	Redis struct {
	} `yaml:"redis"`
	Server struct {
	} `yaml:"server"`
}

func NewConfig() *Config {
	cfg := &Config{}
	err := cleanenv.ReadConfig("config/config.yml", cfg)
	if err != nil {
		log.Fatalln("Dont read config")
		return nil
	}
	return cfg
}
