package cmd

import (
	"errors"
	"fmt"
	"os"

	"reqi/db"

	"github.com/spf13/cobra"
)

var uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "Uninstall request template",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("request template name is required")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		tplName := args[0]
		err := db.DeleteRequestTpl(tplName)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}
