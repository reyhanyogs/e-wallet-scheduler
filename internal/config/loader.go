package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/reyhanyogs/e-wallet-scheduler/internal/component"
)

func Get() *Config {
	err := godotenv.Load("scheduler.env")
	if err != nil {
		component.Log.Fatalf("Get(Load Config): err = %s", err.Error())
	}

	return &Config{
		Redis: Redis{
			Addr: os.Getenv("REDIS_ADDR"),
			Pass: os.Getenv("REDIS_PASS"),
		},
	}
}
