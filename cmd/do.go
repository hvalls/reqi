package cmd

import (
	"errors"
	"fmt"

	"reqi/db"
	"reqi/http"
	"reqi/request"

	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do",
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
			fmt.Println(err)
			return
		}
		r := request.New(tpl, http.NewClient())
		resp, err := r.Execute()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(resp)
	},
}
