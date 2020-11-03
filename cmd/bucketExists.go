package cmd

import (
	"context"
	"log"

	"github.com/spf13/cobra"
)

var (
	bucketToCheck string

	bucketExistsCmd = &cobra.Command{
		Use:   "bucketExists",
		Short: "Check if a bucket exists",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 1 {
				bucketExists(args[0])
			} else {
				log.Fatal("Only one argument is required")
			}
		},
	}
)

func init() {
}

func bucketExists(bucketToCheck string) {
	found, err := MinioClient.BucketExists(context.Background(), bucketToCheck)
	if err != nil {
		log.Fatal(err)
	}
	if found {
		log.Println("Bucket found")
	} else {
		log.Println("Bucket not found")
	}
}
