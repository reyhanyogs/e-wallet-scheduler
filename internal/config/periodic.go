package config

import (
	"os"

	"github.com/hibiken/asynq"
	"github.com/reyhanyogs/e-wallet-scheduler/dto"
	"github.com/reyhanyogs/e-wallet-scheduler/internal/component"
	"gopkg.in/yaml.v3"
)

type PeriodicConfig struct {
	Filename string
}

func (p *PeriodicConfig) GetConfigs() ([]*asynq.PeriodicTaskConfig, error) {
	data, err := os.ReadFile(p.Filename)
	if err != nil {
		component.Log.Errorf("GetConfigs(ReadFile): err = %s", err.Error())
		return nil, err
	}

	var schedules dto.Schedules
	if err = yaml.Unmarshal(data, &schedules); err != nil {
		component.Log.Errorf("GetConfigs(Unmarshal): err = %s", err.Error())
		return nil, err
	}

	var configs []*asynq.PeriodicTaskConfig
	for _, v := range schedules.Config {
		configs = append(configs, &asynq.PeriodicTaskConfig{
			Cronspec: v.Exp,
			Task: asynq.NewTask(
				v.Task,
				nil,
			),
		})
	}

	return configs, nil
}
