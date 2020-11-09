package cmd

import (
	"context"
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/spf13/cobra"
)

var (
	alternativeBucket string

	copyObjectsCmd = &cobra.Command{
		Use:   "copyObjects",
		Short: "Copies the selected object or objects",
		Long:  "Copies the selected object or objects to another path or bucket. Destination could be a path or file.",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 3 {
				copyObjects(args[0], args[1], args[2])
			} else {
				log.Fatal("copyObjects <bucketName> <objectName> <destination>")
			}
		},
	}
)

func init() {
	copyObjectsCmd.Flags().StringVarP(&alternativeBucket, "bucket", "", "", "If given, data is copied to another bucket")
}

func copyObjects(bucket string, object string, destination string) {
	obj, err := MinioClient.GetObject(context.Background(), bucket, object, minio.GetObjectOptions{})
	if err != nil {
		log.Fatal(err)
	}
	statInfo, _ := obj.Stat()
	log.Printf("%+v", statInfo)

	var copyOptions minio.CopyDestOptions
	if alternativeBucket != "" {
		copyOptions = minio.CopyDestOptions{Bucket: alternativeBucket, Object: destination}
	} else {
		copyOptions = minio.CopyDestOptions{Bucket: bucket, Object: destination}
	}
	info, err := MinioClient.CopyObject(context.Background(), copyOptions, minio.CopySrcOptions{Bucket: bucket, Object: object})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Copied:%v from bucket:%v to:%v in bucket:%v", object, bucket, info.Key, info.Bucket)
}
