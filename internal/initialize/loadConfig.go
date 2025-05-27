package initialize

import (
	"fmt"
	"user_service/global"

	"github.com/spf13/viper"
)

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
