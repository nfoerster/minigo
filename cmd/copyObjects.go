package cmd

import (
	"context"
	"log"
	"strings"

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
			} else if len(args) == 1 {
				copyObjects(args[0], ".", ".")
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
	shorthand := false
	var statInfo minio.ObjectInfo

	if object == "/" {
		//shorthand for everything
		shorthand = true
	}
	if !shorthand {
		obj, err := MinioClient.GetObject(context.Background(), bucket, object, minio.GetObjectOptions{})
		if err != nil {
			log.Fatal(err)
		}
		statInfo, _ = obj.Stat()
	}

	if shorthand || (statInfo.Size == 0 && statInfo.ContentType == "") {
		//object is a directory
		if !strings.HasSuffix(object, "/") {
			object = object + "/"
		}
		var objectCh <-chan minio.ObjectInfo
		//folder
		if !shorthand {
			objectCh = MinioClient.ListObjects(context.Background(), bucket, minio.ListObjectsOptions{
				Prefix:    object,
				Recursive: true,
			})
		} else {
			objectCh = MinioClient.ListObjects(context.Background(), bucket, minio.ListObjectsOptions{
				Recursive: true,
			})
		}

		for o := range objectCh {
			if !strings.HasSuffix(destination, "/") {
				destination = destination + "/"
			}
			//recursive call for each object found
			copyObjects(bucket, o.Key, destination)
		}
	} else {
		var copyOptions minio.CopyDestOptions

		if strings.HasSuffix(destination, "/") {
			//place content of s3 object to the desination directory

			filename := strings.Split(object, "/")[len(strings.Split(object, "/"))-1]

			if alternativeBucket != "" {
				copyOptions = minio.CopyDestOptions{Bucket: alternativeBucket, Object: destination + filename}
			} else {
				copyOptions = minio.CopyDestOptions{Bucket: bucket, Object: destination + filename}
			}
			info, err := MinioClient.CopyObject(context.Background(), copyOptions, minio.CopySrcOptions{Bucket: bucket, Object: object})
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("Copied:%v from bucket:%v to:%v in bucket:%v", object, bucket, info.Key, info.Bucket)
		} else {
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
	}

}
