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
		RedisAddr      string `yaml:"redis_addr"`
		RedisDB        string `yaml:"redis_db"`
		RedisDefaultDB string `yaml:"redis_default_db"`
		MinIdleConn    int    `yaml:"min_idle_conn"`
		PoolSize       int    `yaml:"pool_size"`
		PoolTimeout    int    `yaml:"pool_timeout"`
		Password       string `yaml:"password"`
		DB             int    `yaml:"db"`
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
