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
			listBuckets()
		},
	}
)

func init() {
	listBucketsCmd.Flags().BoolVarP(&ts, "ts", "", false, "Print out with creation timestamps")
}

func listBuckets() {
	buckets, err := MinioClient.ListBuckets(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	if ts {
		for _, bucket := range buckets {
			log.Printf("%v [created at %v]", bucket.Name, bucket.CreationDate.Format("2006-01-02 15:04:05"))
		}
	} else {
		for _, bucket := range buckets {
			log.Print(bucket.Name)
		}
	}
}
