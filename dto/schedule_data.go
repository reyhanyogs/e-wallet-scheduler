package dto

type ScheduleData struct {
	Exp  string `yaml:"exp"`
	Task string `yaml:"task"`
}

type Schedules struct {
	Config []ScheduleData `yaml:"configs"`
}
