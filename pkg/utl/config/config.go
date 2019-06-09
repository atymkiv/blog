package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Configuration struct {
	DB     Database `json:"Database, omitempty"`
	JWT    *JWT     `json:"jwt,omitempty"`
	Server *Server  `json:"server,omitempty"`
	Nats   *Nats    `json:"nats,omitempty"`
	Redis  *Redis   `json:"redis,omitempty"`
	GRPC   *GRPC    `json:"grpc,omitempty"`
}

// JWT holds data necessery for JWT configuration
type JWT struct {
	Secret           string `json:"secret,omitempty"`
	Duration         int    `json:"duration_minutes,omitempty"`
	RefreshDuration  int    `json:"refresh_duration_minutes,omitempty"`
	MaxRefresh       int    `json:"max_refresh_minutes,omitempty"`
	SigningAlgorithm string `json:"signing_algorithm,omitempty"`
}

type Database struct {
	Username string
	Password string
	Port     string
	Host     string
	Name     string
}

// Server holds data necessery for server configuration
type Server struct {
	Port         string `json:"port,omitempty"`
	Debug        bool   `json:"debug,omitempty"`
	ReadTimeout  int    `json:"read_timeout_seconds,omitempty"`
	WriteTimeout int    `json:"write_timeout_seconds,omitempty"`
}

type Nats struct {
	QOS  string `json:"qos"`
	Host string `json:"host"`
}

type Redis struct {
	Host     string `json:"host"`
	Password string `json:"password"`
	DB       int    `json:"db"`
}

type GRPC struct {
	ListeningHost string `json:"host"`
}

// Load returns Configuration struct
func Load(path string) (*Configuration, error) {
	bytes, err := ioutil.ReadFile(path)

	if err != nil {
		return nil, fmt.Errorf("error reading config file, %s", err)
	}
	var cfg = new(Configuration)
	err = json.Unmarshal(bytes, cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
