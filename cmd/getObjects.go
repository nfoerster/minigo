package cmd

import (
	"context"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/minio/minio-go/v7"
	"github.com/spf13/cobra"
)

var (
	getObjectsCmd = &cobra.Command{
		Use:   "getObjects",
		Short: "Download the selected objects",
		Long:  "Download the selected objects from either a single selected file or a whole folder",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 3 {
				getObjects(args[0], args[1], args[2])
			} else {
				log.Fatal("getObjects <bucketName> <objectName> <destination>")
			}
		},
	}
)

func getObjects(bucket string, object string, destination string) {
	obj, err := MinioClient.GetObject(context.Background(), bucket, object, minio.GetObjectOptions{})
	if err != nil {
		log.Fatal(err)
	}
	statInfo, _ := obj.Stat()
	if statInfo.Size == 0 && statInfo.ContentType == "" {
		if !strings.HasSuffix(object, "/") {
			object = object + "/"
		}
		//folder
		objectCh := MinioClient.ListObjects(context.Background(), bucket, minio.ListObjectsOptions{
			Prefix:    object,
			Recursive: true,
		})

		for o := range objectCh {
			//recursive call for each object found
			getObjects(bucket, o.Key, destination)
		}

	} else {
		data := make([]byte, statInfo.Size)

		obj.Read(data)

		fi, err := os.Stat(destination)
		if err != nil {
			if strings.Contains(err.Error(), "no such file or directory") {
				//if file does not exist create an empty one
				err := ioutil.WriteFile(destination, []byte{}, 0755)
				if err != nil {
					log.Fatal(err)
				}
				fi, err = os.Stat(destination)
				if err != nil {
					log.Fatal(err)
				}
			} else {
				log.Fatal(err)
			}
		}
		switch mode := fi.Mode(); {
		case mode.IsDir():
			//place content of s3 object to the desination directory

			if !strings.HasSuffix(destination, "/") {
				//path did not end with a slash, add it here
				destination = destination + "/"
			}
			filename := ""
			parts := strings.Split(statInfo.Key, "/")
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
			if _, err := os.Stat(destination); os.IsNotExist(err) {
				//add destination path if not exists
				os.Mkdir(destination, 0755)
			}
			err := ioutil.WriteFile(destination+filename, data, 0755)
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("File:%v from s3 bucket:%v written to:%v", object, bucket, destination+filename)

		case mode.IsRegular():
			//rewrite content to destination file instead of a folder
			err := ioutil.WriteFile(destination, data, 0755)
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("File:%v from s3 bucket:%v written to:%v", statInfo.Key, bucket, destination)
		}
	}
}
