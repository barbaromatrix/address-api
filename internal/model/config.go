package model

type Config struct {
	AppName      string
	ServerPort   string
	HealthPort   string
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
