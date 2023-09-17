package config

import "github.com/spf13/viper"

type config struct {
	ServerHost string
	ServerPort string
	PgHost     string
	PgPort     string
	PgUser     string
	PgPass     string
	PgBase     string
	PgSSLMode  string
}

func ConfInit() (*config, error) {
	viper.SetConfigFile("./../../.env")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	return &config{
		ServerHost: viper.GetString("SERVER_HOST"),
		ServerPort: viper.GetString("SERVER_PORT"),
		PgHost:     viper.GetString("PG_HOST"),
		PgPort:     viper.GetString("PG_PORT"),
		PgUser:     viper.GetString("PG_USER"),
		PgPass:     viper.GetString("PG_PASS"),
		PgBase:     viper.GetString("PG_BASE"),
		PgSSLMode:  viper.GetString("PG_SSL_MODE"),
	}, nil
}
