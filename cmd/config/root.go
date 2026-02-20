package config

import "github.com/spf13/cobra"

var RootCmd = &cobra.Command{
	Use:   "lazycurl",
	Short: "a postman like for terminal",
	Long:  `lazycurl is a command-line tool that provides a Postman-like experience in the terminal.`,
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		panic(err)
	}
}
