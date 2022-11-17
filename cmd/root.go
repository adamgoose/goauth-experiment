package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "goauth",
	Short: "A simple Golang Authentication Experiment",
}

// Execute runs the Cobra Root Command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}

func init() {
	viper.SetEnvPrefix("goauth")
	viper.AutomaticEnv()
}
