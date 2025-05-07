package bar

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthHandler(t *testing.T) {
	tests := []struct {
		name            string
		method          string
		wantStatus      int
		wantResponse    *healthCheckResponse
		wantContentType string
	}{
		{
			name:       "GET request returns 200 and health check response",
			method:     http.MethodGet,
			wantStatus: http.StatusOK,
			wantResponse: &healthCheckResponse{
				Status:  "pass",
				Message: "flux capacitor is fluxing",
			},
			wantContentType: "application/json",
		},
		{
			name:            "OPTIONS request returns 204",
			method:          http.MethodOptions,
			wantStatus:      http.StatusNoContent,
			wantResponse:    nil,
			wantContentType: "",
		},
		{
			name:            "POST request returns 405",
			method:          http.MethodPost,
			wantStatus:      http.StatusMethodNotAllowed,
			wantResponse:    nil,
			wantContentType: "",
		},
		{
			name:            "PUT request returns 405",
			method:          http.MethodPut,
			wantStatus:      http.StatusMethodNotAllowed,
			wantResponse:    nil,
			wantContentType: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// create a request with the test method
			req := httptest.NewRequest(tt.method, "/health", nil)

			// create a ResponseRecorder to record the response
			rr := httptest.NewRecorder()

			// call the handler
			healthHandler(rr, req)

			// check status code
			if status := rr.Code; status != tt.wantStatus {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.wantStatus)
			}

			// check content type header for GET requests
			if tt.wantContentType != "" {
				if contentType := rr.Header().Get("Content-Type"); contentType != tt.wantContentType {
					t.Errorf("handler returned wrong content type: got %v want %v",
						contentType, tt.wantContentType)
				}
			}

			// for GET requests, verify the response body
			if tt.method == http.MethodGet {
				var got healthCheckResponse
				if err := json.NewDecoder(rr.Body).Decode(&got); err != nil {
					t.Fatalf("Failed to decode response body: %v", err)
				}

				if got != *tt.wantResponse {
					t.Errorf("handler returned unexpected body: got %v want %v",
						got, tt.wantResponse)
				}
			}
		})
	}
}
