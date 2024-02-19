package config

type Config struct {
	Redis Redis
}

type Redis struct {
	Addr string
	Pass string
}
