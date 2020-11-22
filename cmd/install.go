package cmd

import (
	"errors"
	"io/ioutil"
	"path/filepath"

	"reqi/db"
	"reqi/requesttpl"

	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install request template",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("request template yaml file is required")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		filename, _ := filepath.Abs(args[0])
		yamlFile, err := ioutil.ReadFile(filename)
		if err != nil {
			panic(err)
		}
		tpl, err := requesttpl.NewYaml(yamlFile)
		if err != nil {
			panic(err)
		}
		err = db.SaveRequestTpl(tpl)
		if err != nil {
			panic(err)
		}
	},
}
