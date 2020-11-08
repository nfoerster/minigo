package cmd

import (
	"log"
	"os"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	cfgSel  string

	endpoint  string
	accesskey string
	secretkey string
	useSSL    bool

	MinioClient *minio.Client

	rootCmd = &cobra.Command{
		Use:   "miniogo",
		Short: "Main cmd",
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
		},
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/minio.yaml)")
	rootCmd.PersistentFlags().StringVar(&cfgSel, "configname", "", "config name in config file, if not given the first is used")

	rootCmd.PersistentFlags().StringVarP(&endpoint, "endpoint", "", "", "The endpoint for the minio connection")
	rootCmd.PersistentFlags().StringVarP(&accesskey, "accesskey", "", "", "The accesskey for minio connection")
	rootCmd.PersistentFlags().StringVarP(&secretkey, "secretkey", "", "", "The secretkey for minio connection")
	rootCmd.PersistentFlags().BoolVarP(&useSSL, "useSSL", "", true, "Connect with ssl")

	rootCmd.AddCommand(createBucketCmd)
	rootCmd.AddCommand(listBucketsCmd)
	rootCmd.AddCommand(bucketExistsCmd)
	rootCmd.AddCommand(bucketToRemoveCmd)
	rootCmd.AddCommand(listObjectsCmd)
	rootCmd.AddCommand(getObjectsCmd)
	rootCmd.AddCommand(removeObjectsCmd)
}

func checkCriticalError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func initConfig() {
	if (accesskey != "" && endpoint == "") ||
		(accesskey == "" && endpoint != "") {
		log.Fatal("If accesskey or endpoint is given, both arguments are required")
	}
	if accesskey != "" {
		if secretkey == "" {
			//only accesskey is given
			secretkey = os.Getenv(accesskey)
			if secretkey == "" {
				log.Fatal("Accesskey is given, but secretkey is not given or is not set in environment variables.")
			}
		}
		connectToMinIO(endpoint, accesskey, secretkey, useSSL)
		return
	}
	//no accesskey is given, use config

	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		checkCriticalError(err)

		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName("minio")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		log.Println("Using config file:", viper.ConfigFileUsed())
	}

	configurations := Configurations{}
	err := viper.Unmarshal(&configurations)
	checkCriticalError(err)

	if cfgSel != "" {
		//select config from config file
		for _, c := range configurations.Configurations {
			{
				if cfgSel == c.Name {
					connectToMinIO(c.Endpoint, c.Accesskey, c.Secretkey, c.UseSSL)
					return
				}
			}
			log.Fatalf("Selected config: %v not found in config file: %v", cfgSel, viper.ConfigFileUsed())
		}
	}
	if len(configurations.Configurations) == 0 {
		log.Fatalf("No valid configurations can be found in config file: %v", viper.ConfigFileUsed())
	}
	connectToMinIO(configurations.Configurations[0].Endpoint, configurations.Configurations[0].Accesskey, configurations.Configurations[0].Secretkey, configurations.Configurations[0].UseSSL)
}

func connectToMinIO(endpoint string, accesskey string, secretkey string, usessl bool) {
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accesskey, secretkey, ""),
		Secure: usessl,
	})
	checkCriticalError(err)
	log.Printf("Successfully connected to: %v", endpoint)
	MinioClient = minioClient
}
