package cmd

import (
	"errors"
	"fmt"

	"reqi/db"
	"reqi/httpclient"
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
			panic(err)
		}
		r := request.New(tpl, httpclient.New())
		resp, err := r.Execute()
		if err != nil {
			panic(err)
		}
		fmt.Println(resp)
	},
}
