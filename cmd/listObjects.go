package cmd

import (
	"context"
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/spf13/cobra"
)

var (
	prefix    string
	recursive bool
	tsList    bool

	listObjectsCmd = &cobra.Command{
		Use:   "listObjects",
		Short: "List objects in a given bucket",
		Run: func(cmd *cobra.Command, args []string) {
			listObjects(args[0], prefix, recursive, tsList)
		},
	}
)

func init() {
	listObjectsCmd.Flags().StringVarP(&prefix, "prefix", "p", "", "Prefix from scratch from")
	listObjectsCmd.Flags().BoolVarP(&recursive, "recursive", "r", false, "Recursive search")
	listObjectsCmd.Flags().BoolVarP(&tsList, "ts", "", false, "Print out with last modification timestamps")
}

func listObjects(name string, prefix string, recursive bool, tsList bool) {

	objectCh := MinioClient.ListObjects(context.Background(), name, minio.ListObjectsOptions{
		Prefix:    prefix,
		Recursive: recursive,
	})

	log.Printf("Content of bucket %v", name)

	for object := range objectCh {
		if object.Err != nil {
			log.Print(object.Err)
			return
		}
		if tsList {
			log.Printf("%v [last modified at %v]", object.Key, object.LastModified.Format("2006-01-02 15:04:05"))
		} else {
			log.Print(object.Key)
		}
	}
}
