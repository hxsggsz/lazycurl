package collections

import (
	"fmt"
	"lazycurl/cmd"
	"lazycurl/cmd/config"
	"os"

	"github.com/spf13/cobra"
)

var collectionName string

var CreateCollectionCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new collection",
	Long:  `Create a new collection in lazycurl.`,
	Run: func(cmd *cobra.Command, args []string) {
		if collectionName == "" {
			panic("Collection name is required")
		}
		ctx := cmd.Context()
		lazycurlPath := ctx.Value(config.LAZYCURL_PATH).(string)

		_, err := os.Stat(lazycurlPath)
		if exists := !os.IsNotExist(err); !exists {

			if err := os.Mkdir(lazycurlPath, os.ModePerm); err != nil {
				panic(err)
			}
		}

		if err := os.Mkdir(lazycurlPath+"/"+collectionName, os.ModePerm); err != nil {
			panic(err)
		}

		fmt.Printf("Collection -> %s created successfully", collectionName)
	},
}

func init() {
	CreateCollectionCmd.Flags().StringVarP(&collectionName, "name", "n", "", "Name of the collection")
	CreateCollectionCmd.MarkFlagRequired("name")
	cmd.RootCmd.AddCommand(CreateCollectionCmd)
}
