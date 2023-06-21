package config

import (
	"flag"
	"log"
	"os"
	"sync"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	AppConfig struct {
		IsDebug   bool   `yaml:"is_debug" env:"APP_IS_DEBUG" env-default:"false"`
		Id        string `yaml:"id" env:"APP_ID"`
		Name      string `yaml:"name" env:"APP_NAME"`
		LogLevel  string `yaml:"log_level" env:"APP_LOG_LEVEL" env-default:"trace"`
		AdminUser struct {
			Email    string `yaml:"email" env:"ADMIN_EMAIL" env-default:"admin"`
			Password string `yaml:"password" env:"ADMIN_PWD" env-default:"admin"`
		} `yaml:"admin"`
	} `yaml:"app"`
	GRPC struct {
		IP   string `yaml:"ip" env:"GRPC_IP"`
		Port int    `yaml:"port" env:"GRPC_PORT"`
	} `yaml:"grpc"`
	HTTP struct {
		IP           string        `yaml:"ip" env:"HTTP_IP"`
		Port         int           `yaml:"port" env:"HTTP_PORT"`
		ReadTimeout  time.Duration `yaml:"read_timeout" env:"HTTP_READ_TIMEOUT"`
		WriteTimeout time.Duration `yaml:"write_timeout" env:"HTTP_WRITE_TIMEOUT"`
		CORS         struct {
			AllowedMethods     []string `yaml:"allowed_methods" env:"HTTP_CORS_ALLOWED_METHODS"`
			AllowedOrigins     []string `yaml:"allowed_origins" env:"HTTP_CORS_ALLOWED_ORIGINS"`
			AllowCredentials   bool     `yaml:"allow_credentials" env:"HTTP_CORS_ALLOW_CREDENTIALS"`
			AllowedHeaders     []string `yaml:"allowed_headers" env:"HTTP_CORS_ALLOWED_HEADERS"`
			OptionsPassthrough bool     `yaml:"options_passthrough" env:"HTTP_CORS_OPTIONS_PASSTHROUGH"`
			ExposedHeaders     []string `yaml:"exposed_headers" env:"HTTP_CORS_EXPOSED_HEADERS"`
			Debug              bool     `yaml:"debug" env:"HTTP_CORS_ALLOWED_DEBUG"`
		} `yaml:"cors"`
	} `yaml:"http"`
	PostgreSQL struct {
		Username string `yaml:"username" env:"PSQL_USERNAME" env-required:"true"`
		Password string `yaml:"password" env:"PSQL_PASSWORD" env-required:"true"`
		Host     string `yaml:"host" env:"PSQL_HOST" env-required:"true"`
		Port     string `yaml:"port" env:"PSQL_PORT" env-required:"true"`
		Database string `yaml:"database" env:"PSQL_DATABASE" env-required:"true"`
	} `yaml:"postgresql"`
}

const (
	EnvConfigPathName  = "CONFIG-PATH"
	FlagConfigPathName = "config"
	HelpText           = "The Art of Development - Production Service"
)

var configPath string
var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		flag.StringVar(&configPath, FlagConfigPathName, "./configs/config.local.yaml", "this is app config file")
		flag.Parse()

		log.Print("config init")

		if configPath == "" {
			configPath = os.Getenv(EnvConfigPathName)
		}

		if configPath == "" {
			log.Fatal("config path is required")
		}

		instance = &Config{}

		if err := cleanenv.ReadConfig(configPath, instance); err != nil {
			helpText := HelpText
			help, _ := cleanenv.GetDescription(instance, &helpText)
			log.Print(help)
			log.Fatal(err)
		}

		if err := cleanenv.ReadEnv(instance); err != nil {
			helpText := HelpText
			help, _ := cleanenv.GetDescription(instance, &helpText)
			log.Print(help)
			log.Fatal(err)
		}
	})
	return instance
}
