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
	ApiGateway struct {
		Http struct {
			Host string
			Port string
		}
	}
	BudgetServer struct {
		Http struct {
			Host string
			Port string
		}
	}
	IncomeAndExpensesServer struct {
		Http struct {
			Host string
			Port string
		}
	}
	UserServer struct {
		Http struct {
			Host string
			Port string
		}
	}
	ReportingServer struct {
		Http struct {
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
