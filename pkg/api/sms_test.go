package api

// Certainly! To test the `SendAfricastalkingBulkSMS` function, you can create a test function using a testing framework like `testing` from the Go standard library. Here's an example test function:

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSendAfricastalkingBulkSMS(t *testing.T) {
	message := "Test message"
	recipients := []string{"+1234567890", "+9876543210"}

	// Mocking HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check the request
		if r.Method != http.MethodPost {
			t.Errorf("Expected POST request, got %s", r.Method)
		}

		if r.URL.String() != "/version1/messaging" {
			t.Errorf("Expected endpoint /version1/messaging, got %s", r.URL.String())
		}

		// Read the request body
		var requestBody map[string]interface{}
		json.NewDecoder(r.Body).Decode(&requestBody)

		// Check the request body
		if requestBody["api_key"] != AfricasTalkingApiKey {
			t.Errorf("Expected API key %s, got %s", AfricasTalkingApiKey, requestBody["api_key"])
		}

		// Check the recipients and message
		batch := requestBody["batch"].([]interface{})
		if len(batch) != len(recipients) {
			t.Errorf("Expected %d recipients in batch, got %d", len(recipients), len(batch))
		}

		for i, recipient := range recipients {
			entry := batch[i].(map[string]interface{})
			if entry["recipient"] != recipient {
				t.Errorf("Expected recipient %s, got %s", recipient, entry["recipient"])
			}

			if entry["message"] != message {
				t.Errorf("Expected message %s, got %s", message, entry["message"])
			}
		}

		// Respond with a mock success response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{"status":"success"}`)
	}))
	defer server.Close()

	// Override the base endpoint with the mock server URL
	BaseSandboxEndpoint = server.URL + "/version1/messaging"

	// Call the function being tested
	err := SendAfricastalkingBulkSMS(message, recipients)

	// Check for errors
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
}

// This test function creates a mock HTTP server using `httptest.NewServer` and checks if the `SendAfricastalkingBulkSMS` function sends the expected request to this mock server. It then verifies the response and checks for any errors. Make sure to replace the placeholder API key with your actual AfricasTalkingApiKey.
