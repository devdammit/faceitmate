package faceit

import (
	"net/http"
	"time"
)

// APIOptions options in constructor
type APIOptions struct {
	ApiKey string
}

type API struct {
	apiKey     string
	hostURI    string
	httpClient http.Client
}

func NewFaceitClient(options APIOptions) *API {
	timeout := 5 * time.Second
	client := http.Client{
		Timeout: timeout,
	}

	return &API{
		apiKey:     options.ApiKey,
		hostURI:    "https://open.faceit.com/data/v4",
		httpClient: client,
	}
}
