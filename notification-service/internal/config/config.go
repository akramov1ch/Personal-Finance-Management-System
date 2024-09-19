package config

import (
	"log"
	"os"

	configloader "github.com/Oyatillohgayratov/config-loader"
)

type Config struct {
	App struct {
		Name        string
		Environment string
	}

	NotifactionServer struct {
		Http struct {
			Host string
			Port string
		}
	}
	
	WebSocketserver struct {
		Http struct {
			Host string
			Port string
		}
	}

	KafkaServer struct {
		Group string
		Topic string
		Http  struct {
			Host string
			Port string
		}
	}
}

func Load() Config {
	cfg := Config{}
	err := configloader.LoadYAMLConfig("config.yaml", &cfg)
	if err != nil {
		log.Fatal("Failed to load configuration", "error", err)
		os.Exit(1)
	}
	return cfg
}
