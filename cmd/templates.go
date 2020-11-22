package cmd

import (
	"os"
	"reqi/db"

	"github.com/kataras/tablewriter"
	"github.com/spf13/cobra"
)

var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List templates",
	Run: func(cmd *cobra.Command, args []string) {

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"name", "description"})

		tpls, err := db.GetRequestTpls()
		if err != nil {
			panic(err)
		}
		for _, tpl := range tpls {
			row := []string{tpl.Name, tpl.Description}
			table.Append(row)
		}
		table.Render()
	},
}
