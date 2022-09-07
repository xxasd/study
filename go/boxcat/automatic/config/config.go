package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

// Create private data struct to hold config options.
type config struct {
	Hostname string `yaml:"hostname"`
	Port     string `yaml:"port"`
}

// Create a new config instance.
var (
	conf *config
)

// New is the constructor of gee.config
func New() *config {
	return &config{
		Hostname: getConf().Hostname,
		Port:     getConf().Port,
	}
}

// Read the config file from the current directory and marshal
// into the conf config struct.
func getConf() *config {
	path, errPath := os.Getwd()
	if errPath != nil {
		panic(errPath)
	}
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()

	if err != nil {
		fmt.Printf("%v", err)
	}

	conf := &config{}
	err = viper.Unmarshal(conf)
	if err != nil {
		fmt.Printf("unable to decode into config struct, %v", err)
	}

	return conf
}
