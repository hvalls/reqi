package cmd

import (
	"errors"
	"fmt"
	"os"

	"reqi/db"
	"reqi/editor"
	"reqi/requesttpl"

	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit request template",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("request template is required")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		tplName := args[0]
		tpl, err := db.GetRequestTpl(tplName)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		tplString, err := tpl.String()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		newYaml, err := editor.EditText(tplString)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		newTpl, err := requesttpl.NewYaml(newYaml)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		err = db.SaveRequestTpl(newTpl)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}
