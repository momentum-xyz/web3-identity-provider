package xhttp

import (
	"encoding/json"
	"net/http"

	"github.com/OdysseyMomentumExperience/web3-identity-provider/pkg/log"
)

type APIError struct {
	Message string `json:"message"`
}

func Error(w http.ResponseWriter, err error, code int) bool {
	if err == nil {
		return false
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)

	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(true)

	// Depending on the error code, log it as either an Error or Debug
	if code == http.StatusInternalServerError {
		log.Logf(0, "%+v", err)
	} else {
		log.Logf(1, "%+v", err)
	}

	err = enc.Encode(APIError{Message: err.Error()})

	if err != nil {
		panic(err) // If this happens, it's a programmer mistake so we panic
	}

	return true
}
