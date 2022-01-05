package go_ambee

import "net/http"

const (
	BaseURLV1 = "https://api.ambeedata.com/latest"
)

type AmbeeClient struct {
	BaseURL string
	apiKey  string
	client  *http.Client
}

type apiResponse struct {
	Message  string    `json:"message"`
	Stations []Station `json:"stations"`
}

func NewClient(apiKey string) *AmbeeClient {
	return &AmbeeClient{
		BaseURL: BaseURLV1,
		apiKey:  apiKey,
		client:  &http.Client{},
	}
}
