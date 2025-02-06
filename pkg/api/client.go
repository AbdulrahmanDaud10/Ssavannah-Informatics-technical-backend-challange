package api

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

var (
	APIKEY              = os.Getenv("AFRICASTALKING_API_KEY_SETTINGS_LABEL")
	USERNAME            = os.Getenv("AFRICASTALKING_USERNAME_SETTINGS_LABEL")
	BASELIVEENDPOINT    = os.Getenv("AFRICASTALKING_BASELIVE_ENDPOINT")
	BASESANDBOXENDPOINT = os.Getenv("AFRICASTALKING_SANDBOX_ENDPOINT")
)

type (
	AtClient struct {
		ApiKey     string `json:"api_key"`
		Endpoint   string `json:"endpoint"`
		httpClient *http.Client
		Username   string `json:"username"`
	}
)

// GetAfricasTalkingSettings returns an instance of an Africa's Talking client reusbale across different products.
func GetAfricasTalkingSettings(apiKey string, username string, URLendPoint string) *AtClient {
	AtClient := &AtClient{
		ApiKey:   apiKey,
		Endpoint: BASESANDBOXENDPOINT,
		httpClient: &http.Client{
			Timeout: time.Second * 10,
		},
		Username: username,
	}

	fmt.Println("##############################", AtClient)
	return AtClient
}

// SetHTTPClient can be used to override the default client with a custom set one.
func (at *AtClient) SetHTTPClient(httpClient *http.Client) *AtClient {
	at.httpClient = httpClient

	return at
}

// setDefaultHeaders sets the standard headers required by the Africa's Talking API.
func (at *AtClient) SetDefaultHeaders(req *http.Request) *http.Request {
	req.Header.Set("apiKey", at.ApiKey)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	return req
}
