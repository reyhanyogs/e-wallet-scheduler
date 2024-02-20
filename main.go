package main

import (
	"time"

	"github.com/hibiken/asynq"
	"github.com/reyhanyogs/e-wallet-scheduler/internal/component"
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
		component.Log.Fatalf("Main(NewPeriodicTaskManager): err = %s", err.Error())
	}

	component.Log.Info("Starting scheduler")
	if err := manager.Run(); err != nil {
		component.Log.Fatalf("Main(Run): err = %s", err.Error())
	}
}
