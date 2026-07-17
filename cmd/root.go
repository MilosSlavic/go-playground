/*
Copyright © 2026 Milos Slavic <m.slavic@outlook.com>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-playground",
	Short: "Playground for Go backend patterns and data structures",
	Long: `go-playground is a collection of small, self-contained Go backend
examples (HTTP server, data structures, CLI/config wiring) used to
demonstrate idiomatic Go patterns.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.go-playground.yaml)")
}
