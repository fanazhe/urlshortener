package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strconv"
)

type Config struct {
	Server struct {
		Host    string `json:"host"`
		Port    string `json:"port"`
		BaseUrl string `json:"baseurl"`
	} `json:"server"`
	Redis struct {
		Host string `json:"host"`
		Port string `json:"port"`
		DB   int    `json:"db"`
	} `json:"redis"`
	APP_ENV string `json:"APP_ENV"`
	WebPath string `json:"webpath"`
}

/*
Load config json file, but ENV variables have higher priority
*/
func Load() (*Config, error) {
	CONFIGFILE, exists := os.LookupEnv("CONFIGFILE")
	if !exists {
		CONFIGFILE = "./config/config.json"
	}
	config, err := FromJson(CONFIGFILE)
	if err != nil {
		return nil, err
	}
	APP_ENV, exists := os.LookupEnv("APP_ENV")
	if exists {
		config.APP_ENV = APP_ENV
	}
	SERVER_HOST, exists := os.LookupEnv("SERVER_HOST")
	if exists {
		config.Server.Host = SERVER_HOST
	}
	SERVER_PORT, exists := os.LookupEnv("SERVER_PORT")
	if exists {
		config.Server.Port = SERVER_PORT
	}
	BACKEND_URL, exists := os.LookupEnv("BACKEND_URL")
	if exists {
		config.Server.BaseUrl = BACKEND_URL
	}
	REDIS_HOST, exists := os.LookupEnv("REDIS_HOST")
	if exists {
		config.Redis.Host = REDIS_HOST
	}
	REDIS_PORT, exists := os.LookupEnv("REDIS_PORT")
	if exists {
		config.Redis.Port = REDIS_PORT
	}
	REDIS_DB, exists := os.LookupEnv("REDIS_DB")
	if exists {
		DBint, err := strconv.Atoi(REDIS_DB)
		if err == nil {
			config.Redis.DB = DBint
		}
	}
	WEBPATH, exists := os.LookupEnv("WEBPATH")
	if exists {
		config.WebPath = WEBPATH
	}
	return config, nil
}

func FromJson(path string) (*Config, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, err
	}
	return &config, nil
}
