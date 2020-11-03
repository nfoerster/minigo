package cmd

type Configuration struct {
	Name      string
	Endpoint  string
	Accesskey string
	Secretkey string
	UseSSL    bool
}

type Configurations struct {
	Configurations []Configuration
}
