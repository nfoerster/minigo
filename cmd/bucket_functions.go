package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var (
	name     string
	location string

	createBucketCmd = &cobra.Command{
		Use:   "create_bucket",
		Short: "Create a new bucket",
		Run: func(cmd *cobra.Command, args []string) {
			log.Print(name)
			log.Print(location)
		},
	}
)

func init() {
	createBucketCmd.Flags().StringVarP(&name, "name", "n", "", "Name of the bucket to create")
	createBucketCmd.MarkFlagRequired("name")
	createBucketCmd.Flags().StringVarP(&location, "location", "l", "us-east-1", "Bucket location")
}
