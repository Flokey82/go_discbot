package go_discbot

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Config represents the configuration of the bot.
type Config struct {
	Token     string `json: "Token"`
	BotPrefix string `json: "BotPrefix"`
}

// ReadConfig reads the configuration file and returns a Config struct.
func ReadConfig(configFile string) (*Config, error) {
	fmt.Println("Reading config file...")
	file, err := ioutil.ReadFile(configFile)

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	fmt.Println(string(file))
	config := &Config{}

	if err = json.Unmarshal(file, &config); err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return config, nil
}
