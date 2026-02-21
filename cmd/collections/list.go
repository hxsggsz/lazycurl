package collections

import (
	"fmt"
	"os"

	"lazycurl/cmd"
	"lazycurl/cmd/config"
	"lazycurl/cmd/utils"

	"github.com/spf13/cobra"
)

var ListCollectionsCmd = &cobra.Command{
	Use:   "list",
	Short: "List all collections",
	Long:  `List all existing collections in lazycurl.`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := cmd.Context()
		lazyCurlPath := ctx.Value(config.LAZYCURL_PATH).(string)

		if exists := utils.FilePathExists(lazyCurlPath); !exists {
			fmt.Println("No collections found. Create one with 'lazycurl create'")
			return
		}

		entries, err := os.ReadDir(lazyCurlPath)
		if err != nil {
			fmt.Printf("Error reading collections directory: %v\n", err)
			os.Exit(1)
		}

		var collections []string
		for _, entry := range entries {
			if entry.IsDir() {
				collections = append(collections, entry.Name())
			}
		}

		if len(collections) == 0 {
			fmt.Println("No collections found.")
			return
		}

		fmt.Printf("Total Collections (%d):\n", len(collections))
		for _, coll := range collections {
			fmt.Printf("  • %s\n", coll)
		}
	},
}

func init() {
	cmd.RootCmd.AddCommand(ListCollectionsCmd)
}
