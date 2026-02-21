package collections

import (
	"fmt"
	"os"
	"path/filepath"

	"lazycurl/cmd"
	"lazycurl/cmd/config"
	"lazycurl/cmd/utils"

	"github.com/spf13/cobra"
)

var collectionNameToDelete string

var DeleteCollectionCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a collection",
	Long:  `Delete an existing collection in lazycurl.`,
	Run: func(cmd *cobra.Command, args []string) {
		if collectionNameToDelete == "" {
			fmt.Println("Collection name is required")
			return
		}

		ctx := cmd.Context()
		lazyCurlPath := ctx.Value(config.LAZYCURL_PATH).(string)
		collectionPath := filepath.Join(lazyCurlPath, collectionNameToDelete)

		if exists := utils.FilePathExists(collectionPath); !exists {
			fmt.Printf("Collection '%s' does not exist\n", collectionNameToDelete)
			return
		}

		if err := os.RemoveAll(collectionPath); err != nil {
			fmt.Printf("Error deleting collection '%s': %v\n", collectionNameToDelete, err)
			return
		}

		fmt.Printf("Collection '%s' deleted successfully\n", collectionNameToDelete)
	},
}

func init() {
	DeleteCollectionCmd.Flags().StringVarP(&collectionNameToDelete, "name", "n", "", "Name of the collection to delete")
	DeleteCollectionCmd.MarkFlagRequired("name")
	cmd.RootCmd.AddCommand(DeleteCollectionCmd)
}
