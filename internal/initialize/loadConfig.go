package initialize

import (
	"fmt"
	"user_service/global"

	"github.com/spf13/viper"
)

/*
@Author: Sinh
@Date: 2025/6/1
@Description: Load configuration from a YAML file using Viper.
The configuration file file is loaded to the global.Config variable.
*/
func LoadConfig() {
	viper := viper.New()
	// Add both the relative path and current directory for flexibility
	viper.AddConfigPath("./config/")
	viper.SetConfigName("dev")
	viper.SetConfigType("yaml")

	// Read in the config file
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	// read the config file
	if err := viper.Unmarshal(&global.Config); err != nil {
		panic(fmt.Errorf("unable to decode configuration into struct, %v", err))
	}
}
