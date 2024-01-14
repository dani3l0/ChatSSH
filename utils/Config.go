package utils

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type config struct {
	ListenAddr    string
	ChatGPTKey    string
	RequireSecret bool
	Secret        string
}

var Config = config{
	ListenAddr:    "0.0.0.0:2222",
	ChatGPTKey:    "openai-key-goes-here",
	RequireSecret: false,
	Secret:        "hackme1337",
}

func LoadConfig() {
	filename := "config.yaml"
	conf, err := os.ReadFile(filename)
	if err != nil {
		conf = []byte("")
		data, _ := yaml.Marshal(Config)
		e := os.WriteFile(filename, data, 0750)
		if e != nil {
			fmt.Println("Couldn't create config file: " + e.Error())
		} else {
			fmt.Println("Config file created")
		}
	}
	yaml.Unmarshal(conf, &Config)
}
