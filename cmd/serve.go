/*
Copyright © 2026 Milos Slavic <m.slavic@outlook.com>
*/
package cmd

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		port := viper.GetInt("port")
		fmt.Printf("Starting server on port: %d\n", port)
		return nil
	},
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		return initilizeConfig(cmd)
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
	serveCmd.Flags().Int("port", 8080, "Port to run the server on")
}

func initilizeConfig(cmd *cobra.Command) error {
	viper.SetEnvPrefix("GOPLG")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "*", "-", "*"))
	viper.AutomaticEnv()

	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)
		viper.AddConfigPath(".")
		viper.AddConfigPath(home + "/.go-playground")
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
	}

	if err := viper.ReadInConfig(); err != nil {
		var fileNotFound viper.ConfigFileNotFoundError
		if !errors.As(err, &fileNotFound) {
			return err
		}
	}

	err := viper.BindPFlags(cmd.Flags())
	if err != nil {
		return err
	}

	fmt.Println("Configuration initialized. Using config file:", viper.ConfigFileUsed())

	return nil
}
