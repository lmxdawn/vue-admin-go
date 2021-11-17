package config

import (
	"github.com/jinzhu/configor"
	"github.com/pkg/errors"
)

func Init(confPath string) error {
	err := initConfig(confPath)
	if err != nil {
		return err
	}
	return nil
}

func initConfig(confPath string) error {
	if confPath != "" {
		err := configor.Load(&Config, confPath)
		if err != nil {
			return errors.WithStack(err)
		}
	} else {
		err := configor.Load(&Config, "config/config.yml")
		if err != nil {
			return errors.WithStack(err)
		}
	}
	return nil
}

type AppConfig struct {
	Port uint `yaml:"port"`
}

type MySQLConfig struct {
	Host            string `yaml:"host"`
	Port            uint   `yaml:"port"`
	Db              string `yaml:"db"`
	UserName        string `yaml:"user_name"`
	Password        string `yaml:"password"`
	MaxIdleConn     int    `yaml:"max_idle_conn"`
	MaxOpenConn     int    `yaml:"max_open_conn"`
	ConnMaxLifeTime int    `yaml:"conn_max_life_time"`
}

// RedisConfig ...
type RedisConfig struct {
	Host         string `yaml:"host"`
	Port         uint   `yaml:"port"`
	Password     string `yaml:"password"`
	Db           int    `yaml:"db"`
	DialTimeout  int    `yaml:"dial_timeout"`
	ReadTimeout  int    `yaml:"read_timeout"`
	WriteTimeout int    `yaml:"write_timeout"`
	PoolSize     int    `yaml:"pool_size"`
}

var Config = struct {
	// common
	App   AppConfig
	MySQL MySQLConfig
	Redis RedisConfig
}{}
