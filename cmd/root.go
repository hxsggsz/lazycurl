package cmd

import (
	"context"
	"fmt"
	"lazycurl/cmd/config"
	"lazycurl/cmd/utils"
	"lazycurl/ui"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "lazycurl",
	Short: "a postman like for terminal",
	Long:  `lazycurl is a command-line tool that provides a Postman-like experience in the terminal.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		collectionName := args[0]

		ctx := cmd.Context()
		lazyCurlPath := ctx.Value(config.LAZYCURL_PATH).(string)
		collectionPath := filepath.Join(lazyCurlPath, collectionName)

		if !utils.FilePathExists(collectionPath) {
			fmt.Printf("Collection '%s' does not exist\n", collectionName)
			fmt.Println("Use 'collections list' to see available collections")
			return
		}

		fmt.Printf("Collection '%s' selected)\n", collectionName)
		ui.InitLayout(collectionPath)
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
