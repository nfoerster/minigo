package cmd

import (
	"context"
	"log"

	"github.com/spf13/cobra"
)

var (
	locationMirroredBucket string

	mirrorBucketCmd = &cobra.Command{
		Use:   "mirrorBucket",
		Short: "Copies a bucket with same content",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 2 {
				mirrorBucket(args[0], args[1])
			} else {
				log.Fatal("copyObjects <sourceBucket> <targetBucket>")
			}
		},
	}
)

func init() {
	mirrorBucketCmd.Flags().StringVarP(&locationMirroredBucket, "location", "l", "us-east-1", "Bucket location")
}

func mirrorBucket(sourceBucket string, targetBucket string) {
	found, err := MinioClient.BucketExists(context.Background(), targetBucket)
	if err != nil {
		log.Fatal(err)
	}
	if !found {
		createBucket(targetBucket, locationMirroredBucket)
	}
	copyAllObjects(sourceBucket, "/", targetBucket)
}
