package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use: "reqi",
}

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(installCmd)
	rootCmd.AddCommand(editCmd)
	rootCmd.AddCommand(uninstallCmd)
	rootCmd.AddCommand(lsCmd)
	rootCmd.AddCommand(doCmd)
}

func Execute() {
	rootCmd.Execute()
}
