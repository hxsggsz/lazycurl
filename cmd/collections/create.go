package collections

import (
	"fmt"
	"lazycurl/cmd"
	"lazycurl/cmd/config"
	"lazycurl/cmd/utils"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var collectionName string

var CreateCollectionCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new collection",
	Long:  `Create a new collection in lazycurl.`,
	Run: func(cmd *cobra.Command, args []string) {
		if collectionName == "" {
			fmt.Printf("Collection name is required")
			return
		}

		ctx := cmd.Context()
		lazyCurlPath := ctx.Value(config.LAZYCURL_PATH).(string)

		if exists := utils.FilePathExists(lazyCurlPath); !exists {
			if err := os.Mkdir(lazyCurlPath, os.ModePerm); err != nil {
				fmt.Printf("Error creating lazyCurlPath directory: %v\n", err)
				return
			}
		}

		if err := os.Mkdir(filepath.Join(lazyCurlPath, collectionName), os.ModePerm); err != nil {
			fmt.Printf("Error creating collection directory: %v\n", err)
			return
		}

		fmt.Printf("Collection -> %s created successfully", collectionName)
	},
}

func init() {
	CreateCollectionCmd.Flags().StringVarP(&collectionName, "name", "n", "", "Name of the collection")
	CreateCollectionCmd.MarkFlagRequired("name")
	cmd.RootCmd.AddCommand(CreateCollectionCmd)
}
