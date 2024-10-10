package config

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		App       `yaml:"app"`
		HTTP      `yaml:"http"`
		Cockroach `yaml:"cockroach"`
		Supabase  `yaml:"supabase"`
		Redis     `yaml:"redis"`
	}

	App struct {
		Name    string `env-required:"true" yaml:"name"    env:"APP_NAME"`
		Version string `env-required:"true" yaml:"version" env:"APP_VERSION"`
	}

	HTTP struct {
		Port string `env-required:"true" yaml:"port" env:"HTTP_PORT"`
	}

	Cockroach struct {
		Url string `env-required:"true" yaml:"url" env:"COCKROACH_URL"`
	}

	Redis struct {
		Addr     string `env-required:"true" yaml:"addr" env:"REDIS_ADDR"`
		Password string `env-required:"true" yaml:"password" env:"REDIS_PASSWORD"`
		Db       string `env-required:"true" yaml:"database" env:"REDIS_DB"`
	}

	Supabase struct {
		Url     string `env-required:"true" yaml:"url" env:"SUPABASE_URL"`
		AnonKey string `env-required:"true" yaml:"anonKey" env:"SUPABASE_ANON_KEY"`
	}
)

func New() *Config {
	cfg := &Config{}

	err := cleanenv.ReadConfig("./configs/config.yaml", cfg)
	if err != nil {
		log.Fatal(err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		log.Fatal(err)
	}

	return cfg
}
