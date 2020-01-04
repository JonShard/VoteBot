package cfg

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

// FileName contains the expected name of the config file to read and write to.
const FileName = "config.yml"

// Config contains the global configuration for the bot.
// Token contains the authorization token for communicating with discord.
type Config struct {
	BotToken         string `yaml:"botToken"`
	Database         string `yaml:"database"`
	DatabaseIP       string `yaml:"databaseIP"`
	DatabaseUser     string `yaml:"databaseUser"`
	DatabasePassword string `yaml:"databasePassword"`
	// Unique per discord server:
	SongLimit       int    `yaml:"songLimit"`
	VotesPerUser    int    `yaml:"votesPerUser"`
	VotesPerPateron int    `yaml:"votesPerPatreon"`
	ChannelID       string `yaml:"textChannelID"`
	MasterRoleID    string `yaml:"masterRoleID"`
}

// Cfg contains the global configuration for the bot.
var Cfg Config

// ReadConfigFile reads config file from file into struct 'config'.
func ReadConfigFile() {
	yamlFile, err := ioutil.ReadFile(FileName)
	if err != nil {
		fmt.Printf("Could not read config file: err #%v ", err)
	}
	Cfg = Config{}
	err = yaml.Unmarshal(yamlFile, &Cfg)
	if err != nil {
		log.Fatalf("Could not unmarshal config file: err #%v", err)
	}

	fmt.Printf("Loaded config file:\n %+v\n", Cfg)
}

// WriteConfigFile saves the current instance of the config struct to file.
func WriteConfigFile() {

	configText, err := yaml.Marshal(Cfg)
	if err != nil {
		fmt.Printf("Could not marshal config struct: err #%v", err)
		return
	}
	err = ioutil.WriteFile(FileName, configText, os.FileMode(int(0777)))
	if err != nil {
		fmt.Printf("Could not write config file: err #%v", err)
		return
	}
}
