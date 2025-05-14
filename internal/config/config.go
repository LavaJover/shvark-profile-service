package config

import (
	"log"
	"os"
	"github.com/ilyakaznacheev/cleanenv"
)

type ProfileConfig struct {
	Env string 	`yaml:"env"`
	GRPCServer 	`yaml:"grpc_server"`
	ProfileDB 	`yaml:"profile_db"`
	LogConfig 	`yaml:"log_config"`
}

type GRPCServer struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type ProfileDB struct {
	Dsn string `yaml:"dsn"`
}

type LogConfig struct {
	LogLevel 	string 	`yaml:"log_level"`
	LogFormat 	string 	`yaml:"log_format"`
	LogOutput 	string 	`yaml:"log_output"`
}

func MustLoad() *ProfileConfig {

	// Processing env config variable and file
	configPath := os.Getenv("PROFILE_CONFIG_PATH")

	if configPath == ""{
		log.Fatalf("PROFILE_CONFIG_PATH was not found\n")
	}

	if _, err := os.Stat(configPath); err != nil{
		log.Fatalf("failed to find config file: %v\n", err)
	}

	// YAML to struct object
	var cfg ProfileConfig
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil{
		log.Fatalf("failed to read config file: %v", err)
	}

	return &cfg
}