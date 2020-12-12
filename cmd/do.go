package cmd

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"reqi/db"
	"reqi/http"
	"reqi/request"

	"github.com/spf13/cobra"
)

func init() {
	doCmd.Flags().StringP("output", "o", "", "output file")
	doCmd.Flags().StringArrayP("parameter", "p", []string{}, "parameter")
}

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
			os.Exit(1)
		}
		r := request.New(tpl, http.NewClient())
		params, err := getParams(cmd)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		resp, err := r.Execute(params)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		writer, err := getWriter(cmd)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Fprintln(writer, resp)
	},
}

func getParams(cmd *cobra.Command) (map[string]string, error) {
	pFlags, err := cmd.Flags().GetStringArray("parameter")
	if err != nil {
		return nil, err
	}
	if len(pFlags) == 0 {
		return map[string]string{}, nil
	}
	params := make(map[string]string)
	for _, p := range pFlags {
		parts := strings.Split(p, "=")
		if len(parts) != 2 {
			return nil, errors.New("invalid parameter")
		}
		params[parts[0]] = parts[1]
	}
	return params, nil
}

func getWriter(cmd *cobra.Command) (*os.File, error) {
	oFlag, err := cmd.Flags().GetString("output")
	if err != nil {
		return nil, err
	}
	if oFlag != "" {
		f, err := os.Create(oFlag)
		if err != nil {
			return nil, err
		}
		return f, nil
	}
	return os.Stdout, nil
}
