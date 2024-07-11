package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"sync"
)

var (
	_mu sync.Mutex
	_c  Config
)

type Config struct {
	Server     Server      `json:"server"`
	AuthServer AuthService `json:"auth_service"`
}

type AuthService struct {
	Enabled bool   `json:"enabled"`
	URl     string `json:"url"`
}

type Server struct {
	DebugMode bool   `json:"debug"`
	Port      int    `json:"port"`
	Host      string `json:"host"`
}

// C return copy of Config
func C() Config {
	return _c
}

func LoadConfig() (*viper.Viper, error) {
	v := viper.New()
	// todo fix
	v.AddConfigPath("src/config")
	v.SetConfigName("config")
	v.SetConfigType("json")
	err := v.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %v", err)
	}
	return v, nil
}

func ParseConfig(v *viper.Viper) (*Config, error) {
	_mu.Lock()
	defer _mu.Unlock()

	err := v.Unmarshal(&_c)
	if err != nil {
		log.Fatalf("unable to decode config into struct, %v", err)
		return nil, err
	}

	return &_c, nil
}
