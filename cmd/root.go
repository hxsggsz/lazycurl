package cmd

import (
	"context"
	"fmt"
	"lazycurl/cmd/config"
	"os"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "lazycurl",
	Short: "a postman like for terminal",
	Long:  `lazycurl is a command-line tool that provides a Postman-like experience in the terminal.`,
	Run: func(cmd *cobra.Command, args []string) {
		// ui.InitLayout()
		fmt.Println("root")
	},
}

func Execute() {
	cfg := config.NewConfig()
	ctx := context.WithValue(context.Background(), config.LAZYCURL_PATH, cfg.LazyCurlPath)
	if err := RootCmd.ExecuteContext(ctx); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
