package model

type IpRequest struct {
	Ip string `json:"ip"`
}

type IpResponse struct {
	Status      string  `json:"status"`
	Message     string  `json:"message"`
	Country     string  `json:"country"`
	CountryCode string  `json:"countryCode"`
	Region      string  `json:"region"`
	State       string  `json:"regionName"`
	City        string  `json:"city"`
	ZipCode     string  `json:"zip"`
	Lat         float64 `json:"lat"`
	Long        float64 `json:"lon"`
	Timezone    string  `json:"timezone"`
}
