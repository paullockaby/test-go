package bar

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type healthCheckResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	if r.Method != http.MethodGet {
		logger.Error(fmt.Sprintf("method %s is not allowed", r.Method))
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// this will always be json, unless the json marshal fails
	w.Header().Set("content-type", "application/json")

	response := healthCheckResponse{
		Status:  "pass",
		Message: "flux capacitor is fluxing",
	}
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		logger.Error(fmt.Sprintf("failed to encode response: %s", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
