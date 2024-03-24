package config

import "github.com/spf13/viper"

type (
	envConfig struct {
		ServerPort string `mapstructure:"SERVER_PORT"`

		DbDriver   string `mapstructure:"DB_DRIVER"`
		DbHost     string `mapstructure:"DB_HOST"`
		DbPort     string `mapstructure:"DB_PORT"`
		DbName     string `mapstructure:"DB_NAME"`
		DbUser     string `mapstructure:"DB_USER"`
		DbPassword string `mapstructure:"DB_PASSWORD"`
	}

	ConfigKey string
)

const (
	ServerPort ConfigKey = "SERVER_PORT"

    DbDriver   ConfigKey = "DB_DRIVER"
    DbHost     ConfigKey = "DB_HOST"
    DbPort     ConfigKey = "DB_PORT"
    DbName     ConfigKey = "DB_NAME"
    DbUser     ConfigKey = "DB_USER"
    DbPassword ConfigKey = "DB_PASSWORD"
)

func LoadEnv(configPath string) error {
	viper.SetConfigType("env")
	viper.AddConfigPath(configPath)
	viper.SetConfigName(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	var env envConfig
	if err := viper.UnmarshalExact(&env); err != nil {
		return err
	}

	return nil
}

func Get(key ConfigKey) string {
	return viper.GetString(string(key))
}
