package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
)

var (
	AfricasTalkingApiKey   = os.Getenv("AFRICASTALKING_API_KEY_SETTINGS_LABEL")
	AfricasTalkingUserName = os.Getenv("AFRICASTALKING_USERNAME_SETTINGS_LABEL")
	BaseLiveEndpoint       = os.Getenv("AFRICASTALKING_BASELIVE_ENDPOINT")
	BaseSandboxEndpoint    = os.Getenv("AFRICASTALKING_SANDBOX_ENDPOINT")
)

type (
	// BulkSMSInput is passed to SendBulkSMS as a parameter.
	BulkSMSInput struct {
		To      []string `json:"to"`
		Message string   `json:"message"`
		From    string   `json:"from"`
	}

	// BulkSMSRecipient is returned as part of the BulkSMSResponse.
	BulkSMSRecipient struct {
		StatusCode uint   `json:"statusCode"`
		Number     string `json:"number"`
		Status     string `json:"status"`
		Cost       string `json:"cost"`
		MessageID  string `json:"messageId"`
	}

	// BulkSMSResponse is returned by SendBulkSMS as a response.
	BulkSMSResponse struct {
		SMSMessageData struct {
			Message    string             `json:"message"`
			Recipients []BulkSMSRecipient `json:"recipients"`
		} `json:"SMSMessageData"`
	}
)

// SendAfricastalkingBulkSMS uses Tuma settings to send a message to multiple phone numbers
func SendAfricastalkingBulkSMS(message string, recipients []string) error {

	bulkMessages := []map[string]string{}
	for _, recipient := range recipients {
		bulkMessages = append(bulkMessages, map[string]string{
			"recipient": recipient,
			"message":   message,
		})
	}

	africastalkingRequestBody := map[string]interface{}{
		"api_key": AfricasTalkingApiKey,
		"batch":   bulkMessages,
	}

	jsonBody, marshallError := json.Marshal(africastalkingRequestBody)

	if marshallError != nil {
		log.Fatal(marshallError)
		return marshallError
	}
	request, httpError := http.NewRequest(http.MethodPost, BaseSandboxEndpoint, bytes.NewBuffer(jsonBody))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("api-key", AfricasTalkingApiKey)

	if httpError != nil {
		log.Fatal(httpError)
		return httpError
	}

	client := http.Client{}

	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	var res map[string]interface{}

	json.NewDecoder(response.Body).Decode(&res)

	if response.StatusCode != http.StatusOK {
		return errors.New(fmt.Sprint(res))
	}
	return nil
}

func SendNotificationSMSAfterOrder() {
}
