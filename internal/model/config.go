package model

type Config struct {
	AppName      string
	ServerPort   string
	HealthPort   string
	TimeLocation string
	ViaCep       HTTPClientConfig
}

type HTTPClientConfig struct {
	HTTP HTTP
}

type HTTP struct {
	URL                 string
	MaxRetry            int
	Timeout             int
	CircuitBreaker      CircuitBreaker
	MaxWaitTimeMillis   int
	MaxIdleConns        int
	MaxIdleConnsPerHost int
}

type CircuitBreaker struct {
	Enabled                 bool
	RequestsInOpenState     uint32
	IntervalMillis          uint32
	OpenStateDurationMillis uint32
	MinRequestsToOpen       uint32
	FailureAllowedRatio     float64
}
