/*
Copyright © 2024 LayerG team
*/
package cmd

import (
	"fmt"
	"layerg-poi-da/global"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var (
	rootCmd = &cobra.Command{
		Use:   "layerg-poi-da",
		Short: "Start a crawler masterdb server",
		Long:  `Start a crawler masterdb server.`,
		Run:   startApi,
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.layerg-masterdb.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find working directory.
		home, err := os.Getwd()
		cobra.CheckErr(err)

		// Search config in working directory with name ".layerg-crawler" (without extension).
		viper.AddConfigPath(home + "/config/")
		viper.SetConfigType("yaml")
		viper.SetConfigName("layerg-masterdb")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}

	// init Config
	if err := viper.Unmarshal(&global.Config); err != nil {
		fmt.Printf("Failed to unmarshal configuration: %v", err)
	}
}
