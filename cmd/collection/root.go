package collection

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "lazycurl",
	Short: "a postman like for terminal",
	Long:  `lazycurl is a command-line tool that provides a Postman-like experience in the terminal.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
