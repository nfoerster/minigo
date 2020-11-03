package cmd

import (
	"context"
	"log"

	"github.com/spf13/cobra"
)

var (
	ts bool

	listBucketsCmd = &cobra.Command{
		Use:   "listBuckets",
		Short: "List all buckets",
		Run: func(cmd *cobra.Command, args []string) {
			listBuckets(name, location)
		},
	}
)

func init() {
	listBucketsCmd.Flags().BoolVarP(&ts, "ts", "", false, "Print our with reation timestamps")
}

func listBuckets(name string, location string) {
	buckets, err := MinioClient.ListBuckets(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	if ts {
		for _, bucket := range buckets {
			log.Printf("%v [created at %v]", bucket.Name, bucket.CreationDate.String())
		}
	} else {
		for _, bucket := range buckets {
			log.Print(bucket.Name)
		}
	}
}
