package config

import (
	"github.com/spf13/viper"
	"path/filepath"
	"strings"
)

type Config struct {
	vip *viper.Viper
}

func New() *Config {
	c := &Config{}
	c.vip = viper.New()
	return c
}

func (c *Config) LoadFileSource(path string) {
	c.vip.AddConfigPath(filepath.Dir(path))
	c.vip.SetConfigName(strings.Replace(filepath.Base(path), filepath.Ext(path), "", -1))
	c.vip.AddConfigPath(".")

}

func (c *Config) Scan(v interface{}) error {
	if err := c.vip.ReadInConfig(); err != nil {
		return err
	}
	return c.vip.Unmarshal(v)
}
