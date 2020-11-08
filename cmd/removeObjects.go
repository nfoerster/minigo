package cmd

import (
	"context"
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/spf13/cobra"
)

var (
	removeObjectsCmd = &cobra.Command{
		Use:   "removeObjects",
		Short: "Removes the given objects",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 2 {
				removeObjects(args[0], args[1])
			} else {
				log.Fatal("removeObjects <bucketName> <objectName>")
			}
		},
	}
)

func removeObjects(bucket string, object string) {

	objectCh := MinioClient.ListObjects(context.Background(), bucket, minio.ListObjectsOptions{
		Prefix:    object,
		Recursive: true,
	})

	for o := range objectCh {
		//recursive call for each object found
		err := MinioClient.RemoveObject(context.Background(), bucket, o.Key, minio.RemoveObjectOptions{})
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("File:%v from s3 bucket:%v removed.", o.Key, bucket)
	}
}
