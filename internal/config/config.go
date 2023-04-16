package config

import (
	"github.com/spf13/viper"
	"log"
)

type Configs struct {
	Server
	Store
}

type Server struct {
	RestPort string `mapstructure:"rest-port"`
}
type Store struct {
	Name      string `mapstructure:"name"`
	Directory string `mapstructure:"directory"`
}

func Init(configFile string) (*Configs, error) {
	viper.SetConfigName(configFile)
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")
	viper.AutomaticEnv()

	if err := parseEnv(); err != nil {
		log.Fatalf("Произошла ошибка при парсинге env файла: %s", err.Error())
		return nil, err
	}

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Произошла ошибка при чтении конфиг файла: %s", err.Error())
		return nil, err
	}

	configs := &Configs{}
	if err := viper.Unmarshal(configs); err != nil {
		log.Fatalf("Произошла ошибка при декодировании структуры файла конфигурации: %s", err.Error())
		return nil, err
	}
	return configs, nil
}

func parseEnv() error {
	configEnvMap := make(map[string]string)

	configEnvMap["server.rest-port"] = "REST_PORT"
	configEnvMap["store.name"] = "STORE_NAME"
	configEnvMap["store.directory"] = "STORE_DIR"

	for configKey, envKey := range configEnvMap {
		if err := viper.BindEnv(configKey, envKey); err != nil {
			return err
		}
	}
	return nil

}
