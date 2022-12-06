package main

import (
	"aurora/cmd"

	"github.com/spf13/viper"
)

func main() {
	viper.SetDefault("DB_HOST", "localhost")
	viper.SetDefault("DB_USER", "root")
	viper.SetDefault("DB_PASSWORD", "root")
	viper.SetDefault("DB_NAME", "aurora")
	viper.SetDefault("WEB_PORT", "3000")

	viper.AutomaticEnv()

	cmd.Execute()
}
