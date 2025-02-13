package initialize

import (
	"fmt"
	"layerg-poi-da/global"

	"github.com/spf13/viper"
)

func LoadConfig() {
	viper := viper.New()

	viper.AddConfigPath("./config/")
	viper.SetConfigName("layerg-masterdb")
	viper.SetConfigType("yaml")

	// read configuration
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	// init Config
	if err := viper.Unmarshal(&global.Config); err != nil {
		fmt.Printf("Failed to unmarshal configuration: %v", err)
	}
}
