package model

type AweSomeRequest struct {
	Cep string `json:"cep"`
}

type AweSomeResponse struct {
	Cep      string `json:"cep"`
	Address  string `json:"address"`
	State    string `json:"state"`
	District string `json:"district"`
	Lat      string `json:"lat"`
	Long     string `json:"lng"`
	City     string `json:"city"`
	Ddd      string `json:"ddd"`
}
