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
				if args[1] == "." || args[1] == "/" {
					copyAllObjects(args[0], args[2], alternativeBucket)
				} else {
					copyObjects(args[0], args[1], args[2], alternativeBucket)
				}
			} else {
				log.Fatal("copyObjects <bucketName> <objectName> <destination>")
			}
		},
	}
)

func init() {
	copyObjectsCmd.Flags().StringVarP(&alternativeBucket, "bucket", "", "", "If given, data is copied to another bucket")
}

func copyObjects(bucket string, object string, destination string, destinationBucket string) {
	var statInfo minio.ObjectInfo

	obj, err := MinioClient.GetObject(context.Background(), bucket, object, minio.GetObjectOptions{})
	if err != nil {
		log.Fatal(err)
	}
	statInfo, _ = obj.Stat()

	if statInfo.Size == 0 && statInfo.ContentType == "" {
		//object is a directory
		if !strings.HasSuffix(object, "/") {
			object = object + "/"
		}
		var objectCh <-chan minio.ObjectInfo
		//folder
		objectCh = MinioClient.ListObjects(context.Background(), bucket, minio.ListObjectsOptions{
			Prefix:    object,
			Recursive: true,
		})

		for o := range objectCh {
			if !strings.HasSuffix(destination, "/") {
				destination = destination + "/"
			}
			//recursive call for each object found
			copyObjects(bucket, o.Key, destination, destinationBucket)
		}
	} else {
		var copyOptions minio.CopyDestOptions

		if destination == "." || destination == "/" {
			//use source name in another bucket
			if destinationBucket == "" {
				log.Fatal("You can't replace a file with itself.")
			}
			copyOptions = minio.CopyDestOptions{Bucket: destinationBucket, Object: object}

		} else if strings.HasSuffix(destination, "/") {
			//place content of s3 object to the desination directory
			var filename string

			parts := strings.Split(object, "/")
			if len(parts) > 1 {
				//ok the object is in a virtual folder on s3
				for i, p := range parts {
					if i == len(parts)-1 {
						//last part is the filename indeed
						filename = p
						break
					}
					//add path parts to destination
					destination = destination + p + "/"
				}
			} else {
				//the object is not in a virtual folder
				filename = parts[0]
			}

			if destinationBucket != "" {
				copyOptions = minio.CopyDestOptions{Bucket: destinationBucket, Object: destination + filename}
			} else {
				copyOptions = minio.CopyDestOptions{Bucket: bucket, Object: destination + filename}
			}
		} else {
			if destinationBucket != "" {
				copyOptions = minio.CopyDestOptions{Bucket: destinationBucket, Object: destination}
			} else {
				copyOptions = minio.CopyDestOptions{Bucket: bucket, Object: destination}
			}
		}
		info, err := MinioClient.CopyObject(context.Background(), copyOptions, minio.CopySrcOptions{Bucket: bucket, Object: object})
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Copied:%v from bucket:%v to:%v in bucket:%v", object, bucket, info.Key, info.Bucket)
	}
}
func copyAllObjects(bucket string, destination string, destinationBucket string) {
	if destinationBucket == "" {
		log.Fatal("Because of recursivity, you can't copy a whole bucket into the same bucket.")
	}
	if destination == "." {
		destination = "/"
	}

	var objectCh <-chan minio.ObjectInfo
	//folder
	objectCh = MinioClient.ListObjects(context.Background(), bucket, minio.ListObjectsOptions{
		Recursive: true,
	})

	for o := range objectCh {
		if !strings.HasSuffix(destination, "/") {
			destination = destination + "/"
		}
		//recursive call for each object found
		copyObjects(bucket, o.Key, destination, destinationBucket)
	}
}
