package cmd

import (
	"context"
	"log"

	"github.com/spf13/cobra"
)

var (
	bucketToRemove string

	bucketToRemoveCmd = &cobra.Command{
		Use:   "removeBucket",
		Short: "Remove if bucket exists",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 1 {
				removeBucket(args[0])
			} else {
				log.Fatal("Only one argument is required")
			}
		},
	}
)

func removeBucket(bucketToRemove string) {
	err := MinioClient.RemoveBucket(context.Background(), bucketToRemove)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Bucket removed")
}
