package config

import (
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
)

// Config

type Config struct {
	Port       uint64
	LogLevel   log.Lvl
	LogOutput  io.Writer
	Production bool
	File       map[string]string
	MongoUrl   string
	SwapiURL   string
}

// Config global
var cfg Config

//GetCfg return global config
func GetCfg() Config {
	return cfg
}

// LoadConfig -
func LoadConfig() Config {
	_ = godotenv.Load(".env")

	var cfg Config

	if s, ok := os.LookupEnv("MONGO_URL"); ok {
		cfg.MongoUrl = s
	}

	if s, ok := os.LookupEnv("API_PORT"); ok {
		v, err := strconv.ParseUint(s, 10, 64)
		if err != nil {
			fmt.Printf("Invalid port %v\n", err)
			panic(err)
		}
		cfg.Port = v
	}

	if s, ok := os.LookupEnv("API_LOG_LEVEL"); ok {
		switch s {
		case "DEBUG":
			cfg.LogLevel = log.DEBUG
		case "INFO":
			cfg.LogLevel = log.INFO
		case "WARN":
			cfg.LogLevel = log.WARN
		case "ERROR":
			cfg.LogLevel = log.ERROR
		case "OFF":
			cfg.LogLevel = log.OFF
		default:
			fmt.Printf("Unknown log level '%s'; using INFO intead.\n", s)
			cfg.LogLevel = log.INFO
		}
	} else {
		cfg.LogLevel = log.INFO
	}

	if s, ok := os.LookupEnv("API_LOG_OUTPUT"); ok {
		switch s {
		case "STDOUT", "":
			cfg.LogOutput = os.Stdout
		case "STDERR":
			cfg.LogOutput = os.Stderr
		default:
			f, err := os.OpenFile(s, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				fmt.Printf("Error setting log output: %v; using STDOUT instead.\n", err)
				cfg.LogOutput = os.Stdout
			} else {
				cfg.LogOutput = f
			}
		}
	} else {
		cfg.LogOutput = os.Stdout
	}

	if s, ok := os.LookupEnv("SWAPI_URL"); ok {
		cfg.SwapiURL = s
	}

	return cfg
}
