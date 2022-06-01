package utils

import "github.com/spf13/viper"

type Config struct {
	MongoURI           string `mapstructure:"MONGODB_URI"`
	DB                 string `mapstructure:"DB"`
	RecipesCol         string `mapstructure:"RECIPES_COLLECTION"`
	UsersCol           string `mapstructure:"USERS_COLLECTION"`
	RedisAddr          string `mapstructure:"REDIS_ADDR"`
	Auth0Domain        string `mapstructure:"AUTH0_DOMAIN"`
	Auth0APIIdentifier string `mapstructure:"AUTH0_API_IDENTIFIER"`
	ServerAddr         string `mapstructure:"SERVER_ADDR"`
}

// LoadConfig reads configuration from environment file or variables
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
