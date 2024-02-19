package main

import (
	"log"
	"time"

	"github.com/hibiken/asynq"
	"github.com/reyhanyogs/e-wallet-scheduler/internal/config"
)

func main() {
	cfg := config.Get()
	redisConnection := asynq.RedisClientOpt{
		Addr:     cfg.Redis.Addr,
		Password: cfg.Redis.Pass,
	}

	periodicConfig := config.PeriodicConfig{
		Filename: "./schedules.yml",
	}

	manager, err := asynq.NewPeriodicTaskManager(asynq.PeriodicTaskManagerOpts{
		RedisConnOpt:               redisConnection,
		PeriodicTaskConfigProvider: &periodicConfig,
		SyncInterval:               5 * time.Second,
	})

	if err != nil {
		log.Fatal("failed to register task: ", err.Error())
	}

	if err := manager.Run(); err != nil {
		log.Fatal("failed to run scheduler: ", err.Error())
	}
}
