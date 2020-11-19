package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Execute request",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("reqi v.0.0.1")
	},
}
