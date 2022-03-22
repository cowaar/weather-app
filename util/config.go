package util

import "github.com/BurntSushi/toml"
import "log"


type Config struct {
	ApiKey string
	FiveDayForecastEndpoint string
}

func ReadConfig() Config {

	var config Config
	if _, err := toml.DecodeFile("config/conf.toml", &config); err != nil {
		log.Fatal(err)
	}
	//log.Print(config.Index)
	return config
}
