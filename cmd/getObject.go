package cmd

import (
	"context"
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/spf13/cobra"
)

var (
	getObjectCmd = &cobra.Command{
		Use:   "getObject",
		Short: "Downloads the object",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 3 {
				getObject(args[0], args[1], args[2])
			} else {
				log.Fatal("getObject <bucketName> <objectName> <destination>")
			}
		},
	}
)

func init() {
	//listBucketsCmd.Flags().BoolVarP(&ts, "ts", "", false, "Print out with creation timestamps")
}

func getObject(bucket string, object string, destination string) {
	obj, err := MinioClient.GetObject(context.Background(), bucket, object, minio.GetObjectOptions{})
	if err != nil {
		log.Fatal(err)
	}

}
