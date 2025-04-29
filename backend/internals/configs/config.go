package configs

import "github.com/spf13/viper"

type Config struct {
	Url            string `mapstructure:"MONGO_URL"`
	Port           int    `mapstructure:"MONGO_PORT"`
	Username       string `mapstructure:"MONGO_ROOT_USERNAME"`
	Password       string `mapstructure:"MONGO_ROOT_PASSWORD"`
	DatabaseName   string `mapstructure:"MONGO_DB_NAME"`
	CollectionName string `mapstructure:"MONGO_COLLECTION_NAME"`
}

func LoadConfig(path string) (*Config, error) {
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
