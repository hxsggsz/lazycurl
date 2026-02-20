package collection

import (
	"fmt"
	"lazycurl/cmd/config"
	"os"

	"github.com/spf13/cobra"
)

var collectionName string

var collectionCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new collection",
	Long:  `Create a new collection in lazycurl.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Creating collection: %s\n", collectionName)
		if collectionName == "" {
			panic("Collection name is required")
		}

		// TODO: put this logic into an struct
		homePath := os.Getenv("HOME") + "/Downloads"
		lazycurlDir := "lazycurl"
		basePath := homePath + "/" + lazycurlDir

		dirs, err := os.ReadDir(homePath)
		if err != nil {
			panic(err)
		}

		// TODO: move this logic when is installing the cli
		found := false
		for _, dir := range dirs {
			dirName := dir.Name()

			if dirName == lazycurlDir {
				found = true
				break
			}
		}

		if !found {
			if err := os.Mkdir(basePath, os.ModePerm); err != nil {
				panic(err)
			}
		}

		if err := os.Mkdir(basePath+"/"+collectionName, os.ModePerm); err != nil {
			panic(err)
		}

		fmt.Printf("Collection -> %s created successfully", collectionName)
	},
}

func init() {
	config.RootCmd.AddCommand(collectionCmd)
	collectionCmd.Flags().StringVarP(&collectionName, "name", "n", "", "Name of the collection")
	collectionCmd.MarkFlagRequired("name")
}
