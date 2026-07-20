package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "smh",
	Short: "smh is quite nothing and nothing, a personal project to manage passwords",
	Long:  "smh - Salted Master Hash is a password management tool designed to be reliably unreliable to the naked eye",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Oops. An error while executing Salted Master Hash '%s'\n", err)
		os.Exit(1)
	}
}
