package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func BindDefault(configPath string) error {
	info, err := os.Stat(configPath)
	if err != nil {
		return err
	}

	if info.IsDir() {
		return fmt.Errorf("%v is not a file", configPath)
	}

	viper.SetConfigFile(configPath)
	viper.SetConfigType("yaml")

	return viper.MergeInConfig()
}
