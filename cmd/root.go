// Package cmd implements helper commands which adds more fun to the application.
//
// Command modules includes 1.Database.
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gin-playground",
	Short: "Happy hacking!",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Happy coding :)")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
