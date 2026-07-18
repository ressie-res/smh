package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var newPassword = &cobra.Command{
	Use:     "np",
	Aliases: []string{"newpassword"},
	Short:   "Create a new password",
	Long:    "Create a new password",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		result := NewPassword(args[0])
		fmt.Printf("%s\n", result)
	},
}

func init() {
	rootCmd.AddCommand(newPassword)
}
