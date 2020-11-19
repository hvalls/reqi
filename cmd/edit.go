package cmd

import (
	"errors"

	"reqi/db"
	"reqi/editor"
	"reqi/requesttpl"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Execute request",
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
			panic(err)
		}
		tplString, err := tpl.String()
		if err != nil {
			panic(err)
		}
		newYaml, err := editor.EditText(tplString)
		if err != nil {
			panic(err)
		}
		var newTpl requesttpl.RequestTpl
		err = yaml.Unmarshal([]byte(newYaml), &newTpl)
		if err != nil {
			panic(err)
		}
		err = db.SaveRequestTpl(newTpl)
		if err != nil {
			panic(err)
		}
	},
}
