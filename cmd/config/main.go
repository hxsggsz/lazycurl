package config

import (
	"encoding/json"
	"fmt"
	"lazycurl/cmd/utils"
	"os"
	"path/filepath"
)

type Config struct {
	configPath   string
	LazyCurlPath string `json:"lazycurl_path"`
}

func NewConfig() *Config {
	configPath := filepath.Join(os.Getenv("HOME"), ".config", "lazycurl")
	lazyCurlPath := filepath.Join(os.Getenv("HOME"), "Downloads", "lazycurl")

	c := &Config{
		configPath:   configPath,
		LazyCurlPath: lazyCurlPath,
	}

	c.initializeConfig()

	return c
}

func (c *Config) initializeConfig() {
	c.createConfigFolder()
	c.createConfigFile()
}

func (c *Config) createConfigFile() {
	if exists := utils.FilePathExists(c.configPath); !exists {
		fmt.Println("Config folder does not exists yet, skipping config.json file.")
		return
	}

	file, err := os.Create(filepath.Join(c.configPath, "config.json"))
	if err != nil {
		fmt.Println("Something went wrong while try to create config file:", err)
		return
	}
	defer file.Close()

	enc := json.NewEncoder(file)
	if err := enc.Encode(c); err != nil {
		fmt.Println("something went wrong while encoding JSON:", err)
	}
}

func (c *Config) createConfigFolder() {
	if exists := utils.FilePathExists(c.configPath); !exists {
		if err := os.Mkdir(c.configPath, os.ModePerm); err != nil {
			panic(err)
		}
	}
}
