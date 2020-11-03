package cmd

import (
	"context"
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/spf13/cobra"
)

var (
	name     string
	location string

	createBucketCmd = &cobra.Command{
		Use:   "createBucket",
		Short: "Create a new bucket",
		Run: func(cmd *cobra.Command, args []string) {
			createBucket(name, location)
		},
	}
)

func init() {
	createBucketCmd.Flags().StringVarP(&name, "name", "n", "", "Name of the bucket to create")
	createBucketCmd.MarkFlagRequired("name")
	createBucketCmd.Flags().StringVarP(&location, "location", "l", "us-east-1", "Bucket location")
}

func createBucket(name string, location string) {
	err := MinioClient.MakeBucket(context.Background(), name, minio.MakeBucketOptions{Region: location})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Successfully created bucket: %v in location: %v", name, location)
}
