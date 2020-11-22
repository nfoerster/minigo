package cmd

import (
	"context"
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/spf13/cobra"
)

var (
	location string

	createBucketCmd = &cobra.Command{
		Use:   "createBucket",
		Short: "Create a new bucket",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 1 {
				createBucket(args[0], location)
			} else {
				log.Fatal("Only one argument is required")
			}
		},
	}
)

func init() {
	createBucketCmd.Flags().StringVarP(&location, "location", "l", "us-east-1", "Bucket location")
}

func createBucket(name string, location string) {
	err := MinioClient.MakeBucket(context.Background(), name, minio.MakeBucketOptions{Region: location})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Successfully created bucket: %v in location: %v", name, location)
}
