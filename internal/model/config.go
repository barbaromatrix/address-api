package model

type Config struct {
	AppName      string
	ServerHost   string
	MetaHost     string
	TimeLocation string
	AweSome      HTTPClientConfig
}

type HTTPClientConfig struct {
	HTTP HTTP
}

type HTTP struct {
	URL             string
	MaxRetry        int
	MaxFailureRatio int
	Name            int
}
