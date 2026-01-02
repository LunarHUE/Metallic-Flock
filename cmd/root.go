package cmd

import (
	"fmt"
	"os"

	"github.com/lunarhue/metallic-flock/cmd/debug"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "metallic",
	Short: "Used to create base k3s cluster.",
	Long:  `Sets up the controller and agent relationship and does cluster authentication.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(debug.RootCmd)
}
