package cmd

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
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
			fmt.Println(err)
			os.Exit(1)
		}
		tpl, err := requesttpl.NewYaml(string(yamlFile))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		err = db.SaveRequestTpl(tpl)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}
